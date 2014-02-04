package main

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

type Cookbook struct {
	Location string
}

type Artifact struct {
	Id       string
	Version  string
	Location string
}

type Release struct {
	Time      int
	Cookbooks []Cookbook
	Artifacts []Artifact
	Deploys   []Deploy
}

type Rule struct {
	validated bool
	message   string
}

type Deploy struct {
	Kind       string
	Artifact   string
	Properties map[string]string
}

type UriType int

const (
	Unknown UriType = 0
	File    UriType = 1
	HTTP    UriType = 2
)

func (release *Release) Display() {
	fmt.Println("time: ", release.Time)
	fmt.Println("artifacts:")
	for _, artifact := range release.Artifacts {
		fmt.Println("\t", artifact.Id, "\t", artifact.Version, "\t", artifact.Location)
	}
	fmt.Println("cookbooks:")
	for _, cookbook := range release.Cookbooks {
		fmt.Println("\t", cookbook.Location)
	}
}

func (release *Release) DisplayLocalDeploy() {
	fmt.Println("#!/bin/sh")
	fmt.Println("export MENU_BASE=`pwd`")
	for _, cookbook := range release.Cookbooks {
		fmt.Println("# Cookbook ", cookbook.Location)
		id := cookbook.Hash()
		fmt.Printf("git clone %s %s\ncd %s\n", cookbook.Location, id, id)
		fmt.Println("vagrant up")
		fmt.Println("cd $MENU_BASE")
	}
	for _, artifact := range release.Artifacts {
		fmt.Println("# Artifact", artifact.Id, artifact.Location)
		kind := getPathType(artifact.Location)
		if kind == File {
			localPath := scrubPath(artifact.Location)
			id := artifact.Hash()
			fmt.Printf("cp %s %s\n", localPath, id)
		}
		if kind == HTTP {
			id := artifact.Hash()
			fmt.Printf("curl %s -o %s\n", artifact.Location, id)
		}
	}
	for _, deploy := range release.Deploys {
		fmt.Println("# Deploy", deploy.Kind, deploy.Artifact, deploy.PropertiesAsString())
		if deploy.Kind == "tomcat" {
			id := Hash([]byte(deploy.Artifact))
			destination := tomcatUrl(deploy)
			fmt.Printf("curl --upload-file %s \"%s\"\n", id, destination)
		}
	}
}

func (release *Release) DisplayArtifacts() {
	for _, artifact := range release.Artifacts {
		fmt.Println(artifact.Id, "\t", artifact.Version, "\t", artifact.Location)
	}
}

func (release *Release) DisplayCookbooks() {
	for _, cookbook := range release.Cookbooks {
		fmt.Println(cookbook.Location)
	}
}

func (release *Release) HasArtifactVersion(version string) bool {
	for _, artifact := range release.Artifacts {
		if artifact.Version == version {
			return true
		}
	}
	return false
}

func (release *Release) HasArtifactVersions(versions []string) bool {
	for _, version := range versions {
		if release.HasArtifactVersion(version) {
			return true
		}
	}
	return false
}

func (release *Release) HasArtifactId(id string) bool {
	for _, artifact := range release.Artifacts {
		if artifact.Id == id {
			return true
		}
	}
	return false
}

func (release *Release) HasArtifactIds(ids []string) bool {
	for _, id := range ids {
		if release.HasArtifactId(id) {
			return true
		}
	}
	return false
}

func (release *Release) HasArtifactLocation(location string) bool {
	for _, artifact := range release.Artifacts {
		if artifact.Location == location {
			return true
		}
	}
	return false
}

func (release *Release) HasArtifactLocations(locations []string) bool {
	for _, location := range locations {
		if release.HasArtifactLocation(location) {
			return true
		}
	}
	return false
}

func (cookbook Cookbook) Hash() string {
	return Hash([]byte(cookbook.Location))
}

func (artifact Artifact) Hash() string {
	return Hash([]byte(artifact.Id))
}

func (deploy Deploy) PropertiesAsString() string {
	props := make([]string, 0)
	for key, value := range deploy.Properties {
		props = append(props, fmt.Sprintf("%s=%s", key, value))
	}
	return strings.Join(props, ", ")
}

func Hash(bytes []byte) string {
	hasher := sha1.New()
	hasher.Write(bytes)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

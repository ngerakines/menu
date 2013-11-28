package main

import (
	"fmt"
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
}

type Rule struct {
	validated bool
	message   string
}

type UriType int

const (
	Unknown UriType = 0
	File    UriType = 1
	S3      UriType = 2
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

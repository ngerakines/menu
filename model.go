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

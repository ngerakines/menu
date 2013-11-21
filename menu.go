package main

import (
	_ "encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/jessevdk/go-flags"
)

type Cookbook struct {
	location string
}

type Artifact struct {
	id string
	version string
	location string
}

type Release struct {
	time      uint32
	cookbooks []Cookbook
	artifacts []Artifact
}

var opts struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	ArtifactIds []string `short:"ai" description:"An artifact id"`
	ArtifactVersions []string `short:"av" description:"An artifact version"`
	ArtifactLocations []string `short:"al" description:"An artifact location"`
}

func help() {
	fmt.Printf(`Usage: menu <command> <arguments and options>

Commands:
	create
	show
	artifacts
	cookbooks
	list
`)	
}

func main() {
	args, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		panic(err)
		help()
		os.Exit(1)
	}
	if len(args) < 2 {
		help()
		os.Exit(1)
	}

	fmt.Printf("Verbosity: %v\n", opts.Verbose)
	fmt.Printf("ArtifactIds: %v\n", opts.ArtifactIds)
	fmt.Printf("ArtifactVersions: %v\n", opts.ArtifactVersions)
	fmt.Printf("ArtifactLocations: %v\n", opts.ArtifactLocations)
	fmt.Printf("Command: %s\n", args[1])
	fmt.Printf("Remaining args: %s\n", strings.Join(args[2:], " "))
}

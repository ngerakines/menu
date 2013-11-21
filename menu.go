package main

import (
	_ "encoding/json"
	"fmt"
	"os"
	_ "strings"
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
	Cookbooks []string `short:"c" description:"A cookbook location"`
}

func help() {
	fmt.Printf(`Usage: menu <command> <arguments and options>

Commands:
	help
	create
	show
	artifacts
	cookbooks
	list
`)
}

func helpHelp() {
	fmt.Printf(`Usage: menu help <command>

Commands:
	help
	create
	show
	artifacts
	cookbooks
	list
`)
}

func helpCreate() {
	fmt.Printf(`Usage: menu create

Options:
	-ai,--artifact-id
	-av,--artifact-version
	-al,--artifact-location
	-c,--cookbook
`)
}

func helpShow() {
	fmt.Printf(`Usage: menu show <path>
`)
}

func helpArtifacts() {
	fmt.Printf(`Usage: menu artifacts <path>
`)
}

func helpCookbooks() {
	fmt.Printf(`Usage: menu cookbooks <path>
`)
}

func helpList() {
	fmt.Printf(`Usage: menu list <options>

Options:
	-ai,--artifact-id
	-av,--artifact-version
	-al,--artifact-location
	-c,--cookbook
`)
}

func main() {
	args, err := flags.Parse(&opts)

	if err != nil {
		panic(err)
		help()
		os.Exit(1)
	}
	if len(args) == 0 {
		help()
		os.Exit(1)
	}

	if args[0] == "help" {
		handleHelp(args)
		os.Exit(0)
	}
	help()

	/*
	fmt.Printf("Verbosity: %v\n", opts.Verbose)
	fmt.Printf("ArtifactIds: %v\n", opts.ArtifactIds)
	fmt.Printf("ArtifactVersions: %v\n", opts.ArtifactVersions)
	fmt.Printf("ArtifactLocations: %v\n", opts.ArtifactLocations)
	fmt.Printf("Cookbooks: %v\n", opts.Cookbooks)
	fmt.Printf("Command: %s\n", args[0])
	fmt.Printf("Remaining args: %s\n", strings.Join(args[1:], " "))
	*/
}

func handleHelp(args []string) {
	if len(args) > 1 {
		if args[1] == "create" {
			helpCreate()
			return
		}
		if args[1] == "show" {
			helpShow()
			return
		}
		if args[1] == "cookbooks" {
			helpCookbooks()
			return
		}
		if args[1] == "artifacts" {
			helpArtifacts()
			return
		}
		if args[1] == "list" {
			helpList()
			return
		}
	}
	helpHelp()
}

func handleCreate(args []string) {
	helpCreate()
}

func handleShow(args []string) {
	helpShow()
}

func handleArtifacts(args []string) {
	helpArtifacts()
}

func handleCookbooks(args []string) {
	helpCookbooks()
}

func handleList(args []string) {
	helpList()
}

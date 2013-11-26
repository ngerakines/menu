package main

import (
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
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
}

var opts struct {
	Verbose           []bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	ArtifactIds       []string `short:"i" long:"artifact-id" description:"An artifact id"`
	ArtifactVersions  []string `short:"v" long:"artifact-version" description:"An artifact version"`
	ArtifactLocations []string `short:"l" long:"artifact-location" description:"An artifact location"`
	Cookbooks         []string `short:"c" long:"cookbook" description:"A cookbook location"`
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
	fmt.Printf(`Usage: menu create <options> [destination]

The create command can be used to create new menu items. A menu item
consists of one or more artifact id, version and location tripples and
one or more cookbooks.

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
	if args[0] == "create" {
		handleCreate(args)
		os.Exit(0)
	}
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
	if len(opts.ArtifactIds) == 0 {
		fmt.Println("Error: One or more artifacts must be provided.")
		fmt.Println()
		helpCreate()
		os.Exit(1)
	}
	if len(opts.ArtifactIds) != len(opts.ArtifactVersions) {
		fmt.Println("Error: An equal number of artifact ids, versions and locations are required.")
		fmt.Println()
		helpCreate()
		os.Exit(1)
	}
	if len(opts.ArtifactVersions) != len(opts.ArtifactLocations) {
		fmt.Println("Error: An equal number of artifact ids, versions and locations are required.")
		fmt.Println()
		helpCreate()
		os.Exit(1)
	}

	artifacts := make([]Artifact, len(opts.ArtifactVersions))
	cookbooks := make([]Cookbook, len(opts.Cookbooks))

	for i := 0; i < len(opts.ArtifactIds); i++ {
		artifact := Artifact{
			Id:       opts.ArtifactIds[i],
			Version:  opts.ArtifactVersions[i],
			Location: opts.ArtifactLocations[i],
		}
		artifacts[i] = artifact
	}

	for i := 0; i < len(opts.Cookbooks); i++ {
		cookbook := Cookbook{
			Location: opts.Cookbooks[i],
		}
		cookbooks[i] = cookbook
	}

	createdAt := time.Now().Unix()

	release := Release{
		Time:      int(createdAt),
		Artifacts: artifacts,
		Cookbooks: cookbooks,
	}

	b, err := json.Marshal(release)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	if len(opts.Verbose) > 0 {
		fmt.Println(string(b))
	}

	path := getCreatePath(release, args)
	dir, fileName := filepath.Split(path)

	if len(opts.Verbose) > 0 {
		fmt.Println("Writing file ", fileName, " to directory ", dir)
	}

	dirExists := false
	dirExists, err = exists(dir)
	if err != nil {
		panic(err)
	}
	if dirExists == false {
		err = os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		panic(err)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getCreatePath(release Release, args []string) string {
	if len(args) > 1 {
		if strings.HasPrefix(args[1], "file:///") {
			return args[1][9:]
		}
		if strings.HasPrefix(args[1], "file://localhost/") {
			return args[1][18:]
		}
		panic("Error: Only file URIs are supported at this time.")
	}
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%v.menu", release.Time)
	return filepath.Join(cwd, fileName)
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

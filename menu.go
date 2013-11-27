package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var opts struct {
	Verbose           []bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	ArtifactIds       []string `short:"i" long:"artifact-id" description:"An artifact id"`
	ArtifactVersions  []string `short:"v" long:"artifact-version" description:"An artifact version"`
	ArtifactLocations []string `short:"l" long:"artifact-location" description:"An artifact location"`
	Cookbooks         []string `short:"c" long:"cookbook" description:"A cookbook location"`
}

func main() {
	args, err := flags.Parse(&opts)

	if err != nil {
		fail(err.Error(), usageHelp)
	}

	switch command := getCommand(args, 1, "usage"); {
	case command == "help":
		handleHelp(args)
	case command == "create":
		handleCreate(args)
	case command == "show":
		handleShow(args)
	case command == "artifacts":
		handleArtifacts(args)
	case command == "cookbooks":
		handleCookbooks(args)
	case command == "list":
		handleList(args)
	default:
		usageHelp()
	}
}

func getCommand(args []string, position int, defaultCommand string) string {
	if len(args) >= position {
		return args[position-1]
	}
	return defaultCommand
}

func handleHelp(args []string) {
	switch command := getCommand(args, 2, "help"); {
	case command == "create":
		helpCreate()
	case command == "show":
		helpShow()
	case command == "cookbooks":
		helpCookbooks()
	case command == "artifacts":
		helpArtifacts()
	case command == "list":
		helpList()
	default:
		helpHelp()
	}
}

func handleCreate(args []string) {
	rules := make([]Rule, 3)
	rules[0] = Rule{len(opts.ArtifactIds) == 0, "Error: One or more artifacts must be provided."}
	rules[1] = Rule{len(opts.ArtifactIds) != len(opts.ArtifactVersions), "Error: An equal number of artifact ids, versions and locations are required."}
	rules[1] = Rule{len(opts.ArtifactVersions) != len(opts.ArtifactLocations), "Error: An equal number of artifact ids, versions and locations are required."}

	for _, rule := range rules {
		if rule.validated {
			fail(rule.message, helpCreate)
		}
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

	path := ""
	path, err = getCreatePath(release, args)
	writeFile(path, b)
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

func getCreatePath(release Release, args []string) (string, error) {
	if len(args) > 1 {
		if strings.HasPrefix(args[1], "file:///") {
			return args[1][9:], nil
		}
		if strings.HasPrefix(args[1], "file://localhost/") {
			return args[1][18:], nil
		}
		return "", errors.New("Error: Only file URIs are supported at this time.")
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("%v.menu", release.Time)
	return filepath.Join(cwd, fileName), nil
}

type DisplayFilter int

func handleShow(args []string) {
	handleDisplay(args, 0, helpShow)
}

func handleArtifacts(args []string) {
	handleDisplay(args, 1, helpArtifacts)
}

func handleCookbooks(args []string) {
	handleDisplay(args, 2, helpCookbooks)
}

func handleDisplay(args []string, filter DisplayFilter, displayFunc help) {
	rules := make([]Rule, 1)
	rules[0] = Rule{len(args) == 1, "Error: One or more paths must be provided."}

	for _, rule := range rules {
		if rule.validated {
			fail(rule.message, displayFunc)
		}
	}

	for index, path := range unique(args) {
		if index > 0 {
			uriType := getPathType(path)
			if uriType == File {
				release, err := readFile(path)
				if err != nil {
					fail(err.Error(), displayFunc)
				}
				switch filter {
				case 0:
					release.Display()
				case 1:
					release.DisplayArtifacts()
				case 2:
					release.DisplayCookbooks()
				}
			}
		}
	}
}

func handleList(args []string) {
	rules := make([]Rule, 1)
	rules[0] = Rule{len(args) == 1, "Error: One or more paths must be provided."}

	for _, rule := range rules {
		if rule.validated {
			fail(rule.message, helpList)
		}
	}

	paths := make([]string, 0)

	for index, path := range unique(args) {
		if index > 0 {
			uriType := getPathType(path)
			if uriType == File {
				filePath := scrubPath(path)
				paths = discoverPaths(filePath, paths)
			}
		}
	}
	for _, path := range paths {
		release, err := readFile(path)
		if err != nil {
			fail(err.Error(), helpList)
		}
		release.Display()
	}
	fmt.Println(paths)
}

func getPathType(path string) UriType {
	if strings.HasPrefix(path, "file:///") {
		return File
	}
	if strings.HasPrefix(path, "file://localhost/") {
		return File
	}
	return Unknown
}

func unique(slice []string) []string {
	values := make([]string, 0)
	for _, element := range slice {
		values = appendIfMissing(values, element)
	}
	return values
}

func appendIfMissing(slice []string, value string) []string {
	for _, ele := range slice {
		if ele == value {
			return slice
		}
	}
	return append(slice, value)
}

package main

import (
	"fmt"
	"os"
)

type help func()

func usageHelp() {
	fmt.Printf(`Usage: menu <command> <arguments and options>

Commands:
	help
	create
	show
	artifacts
	cookbooks
	list
	local-deploy
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

func helpLocalDeploy() {
	fmt.Printf(`Usage: menu local-deploy <path>
`)
}

func displayHelp(reason string, exit bool) {
	fmt.Println(reason)
	fmt.Println()
	helpCreate()
	if exit {
		os.Exit(1)
	}
}

func fail(reason string, displayFunc help) {
	if len(reason) > 0 {
		fmt.Println(reason)
		fmt.Println()
	}
	displayFunc()
	os.Exit(1)
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"encoding/json"
)

func writeFile(path string, b []byte) {
	dir, fileName := filepath.Split(path)

	if len(opts.Verbose) > 0 {
		fmt.Println("Writing file ", fileName, " to directory ", dir)
	}

	dirExists, err := exists(dir)
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

func readFile(path string) (*Release, error) {
	scrubbedPath := scrubPath(path)
	b, err := ioutil.ReadFile(scrubbedPath); if err != nil {
		return nil, err
	}
	var release Release
	err = json.Unmarshal(b, &release)
	if err != nil {
		return nil, err
	}
	return &release, nil
}

func scrubPath(path string) string {
	if strings.HasPrefix(path, "file:///") {
		return path[8:]
	}
	if strings.HasPrefix(path, "file://localhost/") {
		return path[17:]
	}
	return path
}

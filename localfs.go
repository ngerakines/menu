package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
	return nil, nil
}

package main

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

func readFiles() [][2]string {
	dir, err := filepath.Abs(filepath.Dir(""))
	dir = path.Join(dir, "Content")
	TreatError(err)

	files, err := ioutil.ReadDir(dir)
	TreatError(err)

	var contents [][2]string

	for _, file := range files {
		var name = file.Name()
		if !strings.HasSuffix(name, ".txt") {
			continue
		}

		data, err := ioutil.ReadFile(path.Join(dir, name))
		TreatError(err)

		var result = [2]string{name, string(data)}
		contents = append(contents, result)
	}

	return contents
}

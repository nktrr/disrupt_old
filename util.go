package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getAllGoFiles(path string) []string {
	var err error
	goFilesPath := make([]string, 0)
	err = filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			splitName := strings.Split(info.Name(), ".")
			if len(splitName) == 2 && splitName[1] == "go" {
				if err != nil {
					return err
				}
				goFilesPath = append(goFilesPath, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return goFilesPath
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

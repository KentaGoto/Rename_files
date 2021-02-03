package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func main() {
	var dir string
	var fp *os.File

	if len(os.Args) != 3 {
		fmt.Println("The number of arguments specified is incorrect.")
		os.Exit(1)
	} else {
		dir = os.Args[1] // Root directory
		var err error
		fp, err = os.Open(os.Args[2]) // Filename replacement list
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	paths := dirwalk(dir)

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		base := scanner.Text()
		sp := strings.Split(base, "\t")
		src := sp[0] // First column of the replacement list (look for items in this column in the directory)
		dst := sp[1] // Second column of the replacement list (replace with the item in this column)

		for _, path := range paths {
			dirname := filepath.Dir(path)
			filebasename := filepath.Base(path)

			// File name replacement
			if filebasename == src {
				if err := os.Rename(path, dirname+"\\"+dst); err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	fmt.Println("Done!")
}

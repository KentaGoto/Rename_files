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
		dir = os.Args[1] // 起点となるディレクトリ
		var err error
		fp, err = os.Open(os.Args[2]) // ファイル名置換リスト
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
		src := sp[0] // 置換リストの1列目（この列の項目をディレクトリ内で探す）
		dst := sp[1] // 置換リストの2列目（この列の項目に置換する）

		for _, path := range paths {
			dirname := filepath.Dir(path)
			filebasename := filepath.Base(path)

			// ファイル名置換
			if filebasename == src {
				if err := os.Rename(path, dirname+"\\"+dst); err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	fmt.Println("Done!")
}

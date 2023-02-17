package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 这里使用百度的翻译，自己申请一个即可
const API_KEY = "your-api-key"
const SECRET_KEY = "yoursecret"
const PATH = "填入要翻译的项目路径"

func main() {
	run(PATH)
}

func run(sourceDir string) {
	var dirSlice = strings.Split(sourceDir, "/")
	var dirName = dirSlice[len(dirSlice)-1]
	dirSlice = dirSlice[:len(dirSlice)-1]
	var destDir = filepath.Join(strings.Join(dirSlice, "/"), dirName+"-"+"copy")

	createDir(destDir)
	copyFile(sourceDir, destDir)
}

func createDir(dir string) {
	_, err := os.Stat(dir)
	if err == nil {
		return
	}

	if !os.IsNotExist(err) {
		return
	}

	if err := os.Mkdir(dir, 0777); err != nil {
		panic(err)
	}
}

func copyFile(source string, dest string) {
	dirs, err := os.ReadDir(source)
	if err != nil {
		panic(err)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			createDir(filepath.Join(dest, dir.Name()))
			copyFile(filepath.Join(source, dir.Name()), filepath.Join(dest, dir.Name()))
			continue
		}

		fyFile(filepath.Join(source, dir.Name()), filepath.Join(dest, dir.Name()))
	}
	fmt.Println(source, "completed")
}

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var unity map[string]string

var projectPath string

func init() {
	flag.StringVar(&projectPath, "p", ".", "unity project path")
}

func main() {
	flag.Parse()

	if !filepath.IsAbs(projectPath) {
		dir, _ := os.Getwd()
		projectPath = filepath.Join(dir, projectPath)
	}

	success, path, version := FindUnityProject(projectPath)

	if !success {
		fmt.Println("Unity project not found")
		os.Exit(1)
	}

	success, unity = GetInstalledUnity()

	if !success {
		os.Exit(1)
	}

	openProject(path, version)
}

func openProject(dir string, version string) {
	if u, ok := unity[version]; ok {
		err := exec.Command(u, "-projectPath", dir).Start()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Unity " + version + " not found")
	}
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

package main

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
)

func FindUnityProject(basePath string) (bool, string, string) {
	dir := basePath
	for {
		success, version := getUnityVersion(dir)

		if success {
			return true, dir, version
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return false, "", ""
}

func getUnityVersion(dir string) (bool, string) {
	versionFile := path.Join(dir, "ProjectSettings", "ProjectVersion.txt")

	if exists(versionFile) {
		bytes, err := ioutil.ReadFile(versionFile)
		if err != nil {
			panic(err)
		}

		s := string(bytes)
		r := regexp.MustCompile(`m_EditorVersion: (\S+)`)
		m := r.FindStringSubmatch(s)

		if len(m) > 0 {
			return true, m[1]
		}
		return false, ""
	}

	return false, ""
}

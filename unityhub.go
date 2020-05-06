package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func GetInstalledUnity() (bool, map[string]string) {
	hubPath := "C:\\Program Files\\Unity Hub\\Unity Hub.exe"

	if !exists(hubPath) {
		fmt.Println("Unity Hub not found")
		return false, nil
	}

	out, _ := exec.Command(hubPath, "--", "--headless", "e", "-i").CombinedOutput()

	s := string(out)
	r := regexp.MustCompile(`(\S+) , installed at (.+?)[\r\n]`)
	match := r.FindAllStringSubmatch(s, -1)

	unityMap := make(map[string]string)

	for _, v := range match {
		unityMap[v[1]] = v[2]
	}

	return true, unityMap
}

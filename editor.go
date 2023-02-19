package main

import (
	"io/ioutil"
	"strings"
)

func openFile() {
	path := "C:\\Users\\Nekit\\GolandProjects\\GStation\\disrupt.dot"
	buf, err := ioutil.ReadFile(path)
	content := ""
	if err == nil {
		content = string(buf)
	}
	lines := strings.Split(content, ";")
	println(strings.Join(lines, "ENDLINE\n"))
}

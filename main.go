package main

import (
	"bytes"
	"github.com/goccy/go-graphviz"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//path := os.Args[0]
	path := "C:\\Users\\Nekit\\GolandProjects\\avito-test"
	parseProjectToGraph(path)
}

func parseProjectToGraph(path string) {
	goFiles := getAllGoFiles(path)
	files := parseFiles(goFiles)
	visualize(files)
}

func visualize(files []FileInfo) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		for name, _ := range file.structs {
			graph.CreateNode(name)
		}
	}

	var buf bytes.Buffer
	if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
		log.Fatal(err)
	}

	// 2. get as image.Image instance
	_, err = g.RenderImage(graph)
	if err != nil {
		log.Fatal(err)
	}

	// 3. write to file directly
	if err := g.RenderFilename(graph, graphviz.PNG, "C:/Users/Nekit/GolandProjects/disrupt/graph.png"); err != nil {
		log.Fatal(err)
	}

}

func parseFiles(filesPath []string) []FileInfo {
	files := make([]FileInfo, 0)
	for _, path := range filesPath {
		fileInfo, err := parseFile(path)
		if err == nil {
			parseStructs(fileInfo)
			parseFunctions(fileInfo)
		}
		files = append(files, fileInfo)
	}
	return files
}

func parseStructs(file FileInfo) {
	reg := getStruct()
	structs := reg.FindAllString(file.content, -1)
	for _, s := range structs {
		structName := strings.Split(s, " ")[1]
		file.structs[structName] = s
	}
}

func parseFunctions(file FileInfo) {
	reg := getAllFunctions()
	functions := reg.FindString(file.content)
	functionsSplit := strings.Split(functions, "func ")
	for _, s := range functionsSplit {
		println(s)
	}
}

func parseFile(path string) (FileInfo, error) {
	splitPath := strings.Split(path, "\\")
	fileInfo := FileInfo{
		filePath:  path,
		fileName:  splitPath[len(splitPath)-1],
		structs:   make(map[string]string, 0),
		functions: make(map[string]string, 0),
	}
	buf, err := ioutil.ReadFile(path)
	if err == nil {
		fileInfo.content = string(buf)
	}
	return fileInfo, err
}

func getAllGoFiles(path string) []string {
	goFilesPath := make([]string, 0)
	err := filepath.Walk(path,
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

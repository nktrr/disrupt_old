package main

import (
	"bytes"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var path string
	//print("Project directory: ")
	//fmt.Fscan(os.Stdin, &path)
	path = "C:\\Users\\Nekit\\GolandProjects\\GStation"
	parseProject(path)
}

func parseProjectToGraph(path string) {
	goFiles := getAllGoFiles(path)
	files := parseFiles(goFiles)
	findCalls(files, path)
}

func findCalls(files []FileInfo, path string) {
	g := graphviz.New()
	graph, err := g.Graph()

	funcArray := filesToFuncArray(files, graph)
	for _, file := range files {
		for funcName, function := range file.functions {
			funcSignature := file.pack + "." + funcName
			calls := checkCalls(funcArray, function, file.pack)
			for _, call := range calls {
				node1 := funcArray[funcSignature]
				node2 := funcArray[call]
				graph.CreateEdge(call, node1, node2)
			}
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
	pathToPng := path + "\\disrupt.png"
	graph.SetScale(100, 100)
	if err := g.RenderFilename(graph, graphviz.PNG, pathToPng); err != nil {
		log.Fatal(err)
	}
}

func checkCalls(functions map[string]*cgraph.Node, function string, pack string) []string {
	calls := make([]string, 0)
	for s, _ := range functions {
		temp := s
		if strings.Split(temp, ".")[0] == pack {
			temp = strings.Split(temp, ".")[1]
		}
		temp = temp + "("
		if strings.Contains(function, temp) {
			calls = append(calls, s)
		}
	}
	return calls
}

func filesToFuncArray(files []FileInfo, graph *cgraph.Graph) map[string]*cgraph.Node {
	functions := make(map[string]*cgraph.Node, 0)
	for _, file := range files {
		for funcName := range file.functions {
			signature := file.pack + "." + funcName
			functions[signature], _ = graph.CreateNode(signature)
		}
	}
	return functions
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
	functionsSplit := getFuncSignature().Split(functions, -1)
	signatures := getFuncSignature().FindAllString(functions, -1)
	functionsSplit = removeEmptyStrings(functionsSplit)
	for i, s := range functionsSplit {
		s = getCommentary().ReplaceAllString(s, "")
		functionName := getFuncName().FindString(s)
		functionName = signatures[i]
		if simpleFunc().MatchString(functionName) {
			functionName = strings.Split(functionName, "(")[0]
			functionName = strings.Split(functionName, " ")[1]
		} else {
			functionName = strings.Split(functionName, ")")[1]
			functionName = strings.Split(functionName, "(")[0]
			functionName = strings.TrimSpace(functionName)
		}
		s = strings.Replace(s, functionName, "", 1)
		file.functions[functionName] = s
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
	pack := getPackageName().FindString(fileInfo.content)
	pack = strings.Replace(pack, "package ", "", 1)
	fileInfo.pack = pack
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

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

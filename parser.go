package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func parseProject(path string) {
	goFiles := getAllGoFiles(path)
	graph := newGraph()
	files := newParseFiles(goFiles)
	addStructsAndFuncSignatures(graph, files)
	newCheckCalls(graph)
	vizualize(graph, path)
}

func addStructsAndFuncSignatures(graph Graph, files []FileInfo) {
	for _, file := range files {
		fileFunctions := getFunctions(file)
		for _, function := range fileFunctions {
			graph.functions[function.pack+"."+function.name] = function
		}
	}
	print("a")
}

func getFunctions(file FileInfo) []Function {
	funcSigntures := getFuncSignature().FindAllString(file.content, -1)
	funcContent := getFuncSignature().Split(file.content, -1)
	functions := make([]Function, 0)
	for i := 1; i < len(funcSigntures); i++ {
		function := Function{}
		function.pack = file.pack
		function.content = funcContent[i]
		functionName := funcSigntures[i-1]
		function.funcSignature = funcSigntures[i-1]

		// parse function name
		if simpleFunc().MatchString(functionName) {
			functionName = strings.Split(functionName, "(")[0]
			functionName = strings.Split(functionName, " ")[1]
		} else {
			functionName = strings.Split(functionName, ")")[1]
			functionName = strings.Split(functionName, "(")[0]
			functionName = strings.TrimSpace(functionName)
		}
		function.name = functionName

		// check struct
		if regexp.MustCompile("func \\(").MatchString(function.funcSignature) {
			function.funcType = structFunc
		} else {
			function.funcType = nonStructFunc
		}

		// function return parsing
		temp := regexp.MustCompile("[()]").Split(function.funcSignature, -1)
		temp = removeEmptyStrings(temp)

		if function.funcType == structFunc {
			if len(temp) == 3 || len(temp) == 4 {
				function.isReturnType = false
			} else {
				function.returnType = temp[4]
			}
		} else {
			if len(temp) == 2 {
				if nonStructFuncArguments().MatchString(function.funcSignature) {
					function.isReturnType = false
				} else {
					function.isReturnType = true
					function.returnType = temp[1]
				}
			} else if len(temp) == 3 {
				function.isReturnType = true
				function.returnType = temp[2]
			}
		}
		functions = append(functions, function)
	}
	return functions
}

func newCheckCalls(graph Graph) {
	for _, function := range graph.functions {
		for _, possibleFunction := range graph.functions {
			if function.pack != possibleFunction.pack || function.funcSignature != possibleFunction.funcSignature {
				checkFunction(graph, function, possibleFunction)
			}
		}
	}
}

func checkFunction(graph Graph, function Function, checkFunction Function) {
	var reg *regexp.Regexp
	if checkFunction.funcType == structFunc {
		reg = regexp.MustCompile("(go )?[\\w|\\d|.]+" + checkFunction.name)
	} else {
		reg = regexp.MustCompile("(go )?" + checkFunction.name)
	}
	if reg.MatchString(function.content) {
		call := Call{checkFunction, false}
		if strings.Contains("go ", reg.FindString(function.content)) {
			call.goroutine = true
		}
		function.calls = append(function.calls, call)
		graph.functions[function.pack+"."+function.name] = function
	}
}

func newParseFiles(filesPath []string) []FileInfo {
	files := make([]FileInfo, 0)
	for _, path := range filesPath {
		fileInfo, err := newParseFile(path)
		if err == nil {
			files = append(files, fileInfo)
		}
	}
	return files
}

func newParseFile(path string) (FileInfo, error) {
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

func newParseFileFunctions(file FileInfo) {

}

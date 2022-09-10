package main

type Graph struct {
	functions map[string]Function
	structs   []string
}

func newGraph() Graph {
	return Graph{functions: make(map[string]Function), structs: make([]string, 0)}
}

type FileInfo struct {
	fileName  string
	filePath  string
	content   string
	pack      string
	structs   map[string]string
	functions map[string]string
}

type Function struct {
	pack          string
	name          string
	content       string
	funcType      string
	funcStruct    string
	funcSignature string
	isReturnType  bool
	returnType    string
	calls         []Call
}

type Call struct {
	callFunc  Function
	goroutine bool
}

const (
	structFunc    = "objectFunc"
	nonStructFunc = "nonObjectFunc"
	returnFunc    = "returnFunc"
	nonReturnFunc = "nonReturnFunc"
	goroutineCall = "goroutineCall"
)

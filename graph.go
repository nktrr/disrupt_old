package main

type Graph struct {
	functions map[string]Function
	calls     map[Function][]Function
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
	pack string
	name string
}

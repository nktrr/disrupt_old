package main

import "regexp"

func getPackageName() *regexp.Regexp {
	re := regexp.MustCompile("\\Apackage [\\w|\\d]+")
	return re
}

func getAllFunctions() *regexp.Regexp {
	re := regexp.MustCompile("func [\\s|\\S]+")
	return re
}

func getStruct() *regexp.Regexp {
	re := regexp.MustCompile("type [\\S]+ struct {\\n[\\s|\\d|\\w|`:\"]+}")
	return re
}

func getFuncName() *regexp.Regexp {
	re := regexp.MustCompile()
}

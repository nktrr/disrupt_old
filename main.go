package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	print("Project directory: ")
	fmt.Fscan(os.Stdin, &path)
	//path = "C:\\Users\\Nekit\\GolandProjects\\kepler"
	parseProject(path)
}

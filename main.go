package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	print("Project directory: ")
	fmt.Fscan(os.Stdin, &path)
	path = "C:\\1234\\gcache"
	parseProject(path)
}

//func activate(app *gtk.Application) {
//	window := gtk.NewApplicationWindow(app)
//	window.SetTitle("gotk4 Example")
//	window.SetChild(gtk.NewLabel("Hello from Go!"))
//	window.SetDefaultSize(400, 300)
//	window.Show()
//}

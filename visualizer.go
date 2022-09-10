package main

import (
	"bytes"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"log"
)

func vizualize(graph Graph, path string) {
	g := graphviz.New()
	graphv, err := g.Graph()
	nodes := make(map[string]*cgraph.Node, 0)
	for _, function := range graph.functions {
		node, _ := graphv.CreateNode(function.pack + "." + function.name)
		node.SetLabel(function.funcSignature)
		nodes[function.funcSignature] = node
	}

	for _, function := range graph.functions {
		for _, call := range function.calls {
			firstNode := nodes[function.funcSignature]
			secondNode := nodes[call.callFunc.funcSignature]
			graphv.CreateEdge(goroutineCall, firstNode, secondNode)
		}
	}

	var buf bytes.Buffer
	if err := g.Render(graphv, graphviz.PNG, &buf); err != nil {
		log.Fatal(err)
	}
	// 2. get as image.Image instance
	_, err = g.RenderImage(graphv)
	if err != nil {
		log.Fatal(err)
	}

	// 3. write to file directly
	pathToPng := path + "\\disruptnew.png"
	graphv.SetScale(100, 100)
	if err := g.RenderFilename(graphv, graphviz.PNG, pathToPng); err != nil {
		log.Fatal(err)
	}
}
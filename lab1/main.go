package main

import (
	"fmt"
	"github.com/gyuho/goraph"
	"io/ioutil"
	"strconv"
	"strings"
)

// "github.com/gyuho/goraph"
// "gonum.org/v1/gonum/graph"
// "gopkg.in/gyuho/goraph.v2"

func main() {
	edgesNameMap := map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F", 7: "G", 8: "H"}

	graph := goraph.NewGraph()

	bytes, err := ioutil.ReadFile("l1_2.txt")
	if err != nil {
		fmt.Println(err)
	}

	splited := strings.Split(string(bytes), "\n")

	for i, v := range splited {
		if i == 0 {
			fmt.Println(fmt.Sprintf("%s - node count", v))
			continue
		}
		graph.AddNode(goraph.NewNode(edgesNameMap[i]))
	}

	for i, v := range splited {
		if i == 0 {
			continue
		}

		for n, w := range strings.Split(v, " ") {
			if n == i-1 || w == "0" {
				continue
			}

			f, err := strconv.ParseFloat(w, 64)
			if err != nil {
				panic(fmt.Sprintf("Can not parse weigth from '%s' to '%s': %s", edgesNameMap[i], edgesNameMap[n+1], err.Error()))
			}

			graph.AddEdge(goraph.StringID(edgesNameMap[i]), goraph.StringID(edgesNameMap[n+1]), f)
		}

	}

	fmt.Println()
	fmt.Println(graph.String())
	fmt.Println()

	spanningTree, err := goraph.Kruskal(graph)
	if err != nil {
		panic(fmt.Sprintf("Can not execute Kruskal alg: %s", err.Error()))
	}

	fmt.Println("Spanning tree")
	fmt.Println(spanningTree)

	total := 0.0
	for k := range spanningTree {
		total += k.Weight()
	}
	fmt.Println("Total: ", total)
}

/*
	https://github.com/gyuho/goraph
	https://pkg.go.dev
	https://en.wikipedia.org/wiki/Kruskal%27s_algorithm
*/

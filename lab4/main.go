package main

import (
	"fmt"
	"github.com/alonsovidales/go_graph"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("l4-1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	edges := make([]graphs.Edge, 0)
	for i, v := range strings.Split(string(bytes), "\n") {
		if i == 0 {
			continue
		}

		for i2, v2 := range strings.Split(v, " ") {

			n, err := strconv.ParseFloat(v2, 32)
			if err != nil {
				panic(err)
			}

			edges = append(edges, graphs.Edge{
				From:   uint64(i),
				To:     uint64(i2 + 1),
				Weight: n,
			})
		}
	}
	graph2 := graphs.GetGraph(edges, false)
	a, b, c := graph2.MinCutMaxFlow(uint64(1), uint64(8), 1.0)
	fmt.Println("Max flow min count: ", a)
	fmt.Println("Flows: ", b)
	fmt.Println()
	tw := 0.0
	for _, v := range c {
		fmt.Println(fmt.Sprintf("From: %d to %d waigth: %f", v.From, v.To, v.Weight))
		tw += v.Weight
	}

	fmt.Println("Total weight: ", tw)
}

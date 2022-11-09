package main

import (
	"fmt"
	"germainewongg/DSA/graphs"
)

func main() {
	graph_display()
}

func graph_display() {
	fmt.Println("Initialising a graph")
	test := graphs.New()
	for i := 0; i < 6; i++ {
		test.AddNode(i)
	}

	// We add some edges
	test.AddEdge(0, 4)
	test.AddEdge(0, 2)
	test.AddEdge(0, 1)
	test.AddEdge(4, 5)
	test.Print()

	// We visit each node through bfs. Visits displays the order through which we visited the node
	fmt.Println("BFS Traversal")
	visitsBFS := test.BFS(0)
	fmt.Println(visitsBFS)

	fmt.Println("DFS Traversal")
	var visitedNodes []int
	visitsDFS := test.DFS(0, visitedNodes)
	fmt.Println(visitsDFS)
}

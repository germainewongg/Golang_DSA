package graphs

import "fmt"

type Graph struct {
	Nodes []*Node
}

type Node struct {
	key      int
	adjacent []*Node
}

func New() *Graph {
	return &Graph{}
}

// Prints out the node connections
func (g *Graph) Print() {
	for _, v := range g.Nodes {
		fmt.Printf("Node %v:", v.key)
		for _, i := range v.adjacent {
			fmt.Printf(" %v", i.key)
		}
		fmt.Println()
	}
}

// We check if the node already exists in the graph
func Contains(s []*Node, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Adds nodes to the graph
func (g *Graph) AddNode(k int) {
	existing := Contains(g.Nodes, k)

	// If node already exists we exit.
	if existing {
		fmt.Println("Node already exists")
	} else {
		g.Nodes = append(g.Nodes, &Node{key: k})
	}
}

// We connect the nodes in the graph
func (g *Graph) AddEdge(n1, n2 int) {
	// Get nodes
	fromNode := g.getNode(n1)
	toNode := g.getNode(n2)

	if fromNode == nil || toNode == nil {
		fmt.Println("Invalid insertion")
	} else {
		existing := Contains(fromNode.adjacent, toNode.key)

		if !existing {
			fromNode.adjacent = append(fromNode.adjacent, toNode)
		} else {
			fmt.Println("Nodes are already connected")
		}

	}

}

// Returns the address of nodes for connection
func (g *Graph) getNode(k int) *Node {

	for i, v := range g.Nodes {
		if v.key == k {
			return g.Nodes[i]
		}
	}

	return nil
}

package graphs

// Recursive DFS implementation
func (g *Graph) DFS(startNode int, visitedNodes []int) []int {
	visited := map[int]bool{startNode: true}

	if !contains(visitedNodes, startNode) {
		visitedNodes = append(visitedNodes, startNode)
	}

	for _, v := range g.Nodes[startNode].adjacent {
		if visited[v.key] {
			continue
		} else {
			visitedNodes = append(visitedNodes, v.key)
			visitedNodes = g.DFS(v.key, visitedNodes)

		}
	}
	return visitedNodes
}

// Checks if the node has already been added to the visited slice
func contains(search []int, key int) bool {
	if len(search) == 0 {
		return false
	}

	for _, v := range search {
		if key == v {
			return true
		}
	}

	return false

}

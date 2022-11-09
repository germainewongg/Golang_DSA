package graphs

// Search for endpoint and return end node
func (g *Graph) BFS(startNode int) (orderedVisits []int) {
	visited := map[int]bool{}

	queue := []int{startNode}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		orderedVisits = append(orderedVisits, current)

		// Start traversal
		for _, neighbour := range g.Nodes[startNode].adjacent {
			if visited[neighbour.key] {
				continue
			}

			visited[neighbour.key] = true
			queue = append(queue, neighbour.key) // Enqueue to continue traversal
		}
	}
	return orderedVisits
}

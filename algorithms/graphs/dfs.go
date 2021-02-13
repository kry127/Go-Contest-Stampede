package graphs

/// DFS WITH CALLBACKS
// author: kry127

// time complexity: O(n)

func DFS(startVertex int, graph map[int][]int, inCallback, outCallback func(vertex, time int)) {
	time := 0
	visited := make(map[int]bool)

	var innerDFS func(v int)
	innerDFS = func (v int) {
		visited[v] = true
		inCallback(v, time)
		time++

		verts, any := graph[v]
		if any {
			for _, w := range verts {
				_, wVisited := visited[w]
				if !wVisited {
					visited[w] = true
					innerDFS(w)
				}
			}
		}

		outCallback(v, time)
		time++
	}
	innerDFS(startVertex)
}
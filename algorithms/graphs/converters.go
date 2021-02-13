package graphs

func EdgesToBidirectGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, e := range edges {
		x, y := e[0], e[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}
	return graph
}

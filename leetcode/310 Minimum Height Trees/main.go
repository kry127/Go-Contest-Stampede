package main

import "fmt"

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

func EdgesToBidirectGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, e := range edges {
		x, y := e[0], e[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}
	return graph
}

func findMinHeightTrees(n int, edges [][]int) []int {
	graph := EdgesToBidirectGraph(edges)

	// first DFS to find deepest verticle
	height := -1
	maxHeight := -1
	deepestX := -1
	DFS(0, graph,
		func(vertex, time int) {
			height++
			if height > maxHeight {
				deepestX = vertex
				maxHeight = height
			}
		},
		func(vertex, time int) {
			height--
		},
	)

	// second DFS to find other deepest verticle
	height = -1
	maxHeight = -1
	deepestY := -1
	lastBacktrackLvl := -1
	var backtrack []int
	DFS(deepestX, graph,
		func(vertex, time int) {
			height++
			if height > maxHeight {
				deepestY = vertex
				maxHeight = height
				lastBacktrackLvl = height
				backtrack = []int{vertex}
			}
		},
		func(vertex, time int) {
			if height < lastBacktrackLvl {
				backtrack = append(backtrack, vertex)
				lastBacktrackLvl = height
			}
			height--
		},
	)

	// then take middle of the backtrack as an answer
	mid := len(backtrack) / 2
	if len(backtrack) % 2 == 0 {
		return  backtrack[mid-1:mid + 1]
	} else {
		return backtrack[mid:mid + 1]
	}
}

func main() {
	//res := findMinHeightTrees(6, [][]int{{3,0},{3,1},{3,2},{3,4},{5,4}})
	res := findMinHeightTrees(6, [][]int{})
	fmt.Printf("%v\n", res)
}

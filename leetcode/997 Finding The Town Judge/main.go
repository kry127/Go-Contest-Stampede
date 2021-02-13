package main

func findJudge(N int, trust [][]int) int {
	indegree := make([]int, N)
	outdegree := make([]int, N)
	for _, v := range trust {
		a, b := v[0] - 1, v[1] - 1
		outdegree[a]++
		indegree[b]++
	}

	for i := range indegree {
		in := indegree[i]
		out := outdegree[i]
		if in == N - 1 && out == 0 {
			return i + 1
		}
	}
	return -1
}

func main() {

}

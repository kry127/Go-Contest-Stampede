package main

import "fmt"

var cache = make(map[int][]string)
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	result, cached := cache[n]
	if cached {
		return result
	}

	result = make([]string, 0)
	for i := 0; i < n; i++ {
		j := n - 1 - i
		// (<inner>)<outer>
		inner := generateParenthesis(j)
		outer := generateParenthesis(i)
		row := make([]string, len(inner) * len(outer))
		k := 0
		for _, in := range inner {
			for _, out := range outer {
				row[k] = "(" + in + ")" + out
				k++
			}
		}
		result = append(result, row...)
	}
	cache[n] = result
	return result
}

func main() {
	res := generateParenthesis(8)
	fmt.Printf("%v\n", res)
}

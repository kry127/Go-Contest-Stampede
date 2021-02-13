package main

import "fmt"

func eliminate(stack []int, v int) ([]int, bool) {
	if len(stack) == 0 {
		return stack, false
	}
	stack, w := stack[:len(stack) - 1], stack[len(stack) - 1]
	switch {
	case w == '(' && v == ')': return stack, true
	case w == '{' && v == '}': return stack, true
	case w == '[' && v == ']': return stack, true
	}
	return stack, false
}
func isValid(s string) bool {
	stack := make([]int, 0)
	for _, x := range s {
		v := int(x)
		ok := true
		switch v {
		case '(' : stack = append(stack, v)
		case '{' : stack = append(stack, v)
		case '[' : stack = append(stack, v)
		case ')' : stack, ok = eliminate(stack, v)
		case '}' : stack, ok = eliminate(stack, v)
		case ']' : stack, ok = eliminate(stack, v)
		}
		if !ok {
			return false
		}
	}
	return true
}

func main() {
	res := isValid("()[]{}")
	fmt.Printf("%v\n", res)
}

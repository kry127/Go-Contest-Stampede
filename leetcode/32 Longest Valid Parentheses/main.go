package main

import "fmt"

type StackElem struct {
	code int32 // '(', ')', '*'
	len uint32
}

func pushStar(stack []StackElem, star StackElem, star_callback func (elem StackElem)) []StackElem {
	star_callback(star) // notify new star pushed to stack
	if len(stack) == 0 {
		return append(stack, star)
	}
	value := stack[len(stack)-1]
	switch value.code {
	case '(': return append(stack, star)
	case '*': return pushStar(stack[:len(stack)-1], StackElem{'*', value.len + star.len}, star_callback)
	}
	panic("Invalid stack state")
}

func pushBracket(stack []StackElem, bracket StackElem, star_callback func (elem StackElem)) []StackElem {
	if len(stack) == 0 {
		return stack
	}
	stack, value := stack[:len(stack)-1], stack[len(stack)-1]
	switch value.code {
	case '(': return pushStar(stack, StackElem{'*', 1 + bracket.len}, star_callback)
	case '*': return pushBracket(stack, StackElem{')', value.len + bracket.len}, star_callback)
	}
	panic("Invalid stack state")
}

func longestValidParentheses(s string) int {
	stack := make([]StackElem, 0)
	max_len := uint32(0)
	maxifier := func (star StackElem) {
		if star.len > max_len {
			max_len = star.len
		}
	}
	for _, ch := range s {
		switch ch {
		case '(': stack = append(stack, StackElem{ch, 1})
		case ')': stack = pushBracket(stack, StackElem{ch, 1}, maxifier)
		}
	}
	return int(max_len)
}

func main() {
	res := longestValidParentheses("")
	fmt.Printf("%d\n", res)
}

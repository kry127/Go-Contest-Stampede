package main

import "fmt"

type Invalids struct {
	left, right int
}

func calcInvalids(s string) Invalids {
	ret := Invalids{}
	for _, v := range s {
		switch v {
		case '(': ret.left += 1
		case ')':
			if ret.left == 0 {
				ret.right++
			} else {
				ret.left--
			}
		}
	}
	return ret
}

func removeInvalidParenthesesRec(s string, inv Invalids) []string {
	if inv.left == 0 && inv.right == 0 {
		return []string{s}
	}

	ret := make([]string, 0)
	for i, v := range s {
		switch {
		case v == '(' && inv.left > 0:
			substrings := removeInvalidParenthesesRec(s[i+1:], Invalids{inv.left-1, inv.right})
			for k := range substrings {
				substrings[k] = s[:i] + substrings[k]
			}
			ret = append(ret, substrings...)
		case v == ')' && inv.right > 0:
			substrings := removeInvalidParenthesesRec(s[i+1:], Invalids{inv.left, inv.right - 1})
			for k := range substrings {
				substrings[k] = s[:i] + substrings[k]
			}
			ret = append(ret, substrings...)
		}
	}
	return ret
}

func removeInvalidParentheses(s string) []string {
	resultsToFilter := removeInvalidParenthesesRec(s, calcInvalids(s))
	uniqResultsMap := make(map[string]bool)
	for _, rtf := range resultsToFilter {
		inv := calcInvalids(rtf)
		if inv.left == 0 && inv.right == 0 {
			uniqResultsMap[rtf] = true
		}
	}

	uniqResults := make([]string, len(uniqResultsMap))
	i := 0
	for rtf := range uniqResultsMap {
		uniqResults[i] = rtf
		i++
	}
	return uniqResults
}

func main() {
	answ := removeInvalidParentheses("(a)())()")
	fmt.Printf("%v\n", answ)
}

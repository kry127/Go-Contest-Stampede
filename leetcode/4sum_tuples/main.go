package main

import (
	"fmt"
	"sort"
)

type Quadruple struct {
	a, b, c, d int
}

type Tuple struct {
	a, b int // values
}

func counter(nums [] int) map[int]int {
	// count nums
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	return count
}

func fourSum(nums []int, target int) [][]int {
	// count nums
	count := counter(nums)

	// make slice of tuples
	uniqTuples := make(map[Tuple]bool)
	tuples := make([]Tuple, 0)
	for i, v := range nums {
		for _, w := range nums[i+1:] {
			tupl := Tuple{v, w}
			_, present := uniqTuples[tupl]
			if !present {
				tuples = append(tuples, Tuple{v, w})
				uniqTuples[tupl] = true
			}
		}
	}

	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].a + tuples[i].b < tuples[j].a + tuples[j].b
	}) // inplace

	// for quadruples uniqueness
	strangeMap := make(map[Quadruple]bool)
	res := make([][]int, 0)
	left := 0
	right := len(tuples) - 1
	for left < len(tuples) && right >= 0 {
		s := tuples[left].a + tuples[left].b + tuples[right].a + tuples[right].b
		if s > target {
			right--
		} else if s < target {
			left++
		} else {
			ll := left
			for tuples[ll].a + tuples[ll].b == tuples[left].a + tuples[left].b {
				ll++
				if ll == len(tuples) {
					break
				}
			}
			ll--
			rr := right
			for tuples[rr].a + tuples[rr].b == tuples[right].a + tuples[right].b {
				rr--
				if rr < 0 {
					break
				}
			}
			rr++
			// we found square [left..ll] x [rr..right]

			for _, tupleL := range tuples[left:ll+1] {
				for _, tupleR := range tuples[rr:right+1] {

					subslice := []int{tupleL.a, tupleL.b, tupleR.a, tupleR.b}
					sort.Ints(subslice)
					sscount := counter(subslice)

					u, v, w, x := subslice[0], subslice[1], subslice[2], subslice[3]

					// check solution validity
					valid := true
					for k, v := range sscount {
						if count[k] < v {
							valid = false
							break
						}
					}
					if !valid {
						continue
					}

					qdrpl := Quadruple{u, v, w, x}
					_, exists := strangeMap[qdrpl]
					if !exists {
						res = append(res, []int{u, v, w, x})
						strangeMap[qdrpl] = true
					}
				}
			}

			left = ll + 1
			right = rr - 1
		}
	}
	return res
}

func main() {
	res := fourSum([]int{3, -4, 5, -6, 7, 2, 2, 3, 3, 4, 4, 5, 5, 5}, 2)
	//res := fourSum([]int{0, 0, 0, 0}, 0)
	fmt.Printf("%v\n", res)
}

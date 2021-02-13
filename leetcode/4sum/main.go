package main

import (
	"fmt"
	"sort"
)

func binsearch(array []int, target int) int {
	return binsearch_ext(array, target, 0, len(array))
}

func binsearch_ext(array []int, target int, startIndex int, endIndex int) int {
	var mid int
	for startIndex < endIndex {
		mid = (endIndex + startIndex)/2
		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid+1
		} else {
			return mid
		}
	}
	return -1
}

type Quadruple struct {
	a, b, c, d int
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums) // inplace
	strangeMap := make(map[Quadruple]bool)
	res := make([][]int, 0)
	for i, u := range nums {
		for j, v := range nums[i+1:] {
			j += i + 1
			for k, w := range nums[j+1:] {
				k += j + 1
				x := target - u - v - w
				xid := binsearch_ext(nums, x, k + 1, len(nums))
				if xid == -1 {
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
	}
	return res
}

func main() {
	//res := fourSum([]int{3, -4, 5, -6, 7, 2, 2, 3, 3, 4, 4, 5, 5, 5}, 2)
	res := fourSum([]int{0, 0, 0, 0}, 0)
	fmt.Printf("%v\n", res)
}

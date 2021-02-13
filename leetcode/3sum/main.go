package main

import "fmt"

func threeSum(nums []int) [][]int {
	numToCount := make(map[int]int)
	for _, num := range nums {
		numToCount[num]++
	}
	fmt.Printf("%v\n", numToCount)

	ret := make([][]int, 0)
	for key1, key1cnt := range numToCount {
		for key2 := range numToCount {
			key3 := -(key1 + key2)
			key3cnt, key3Exists := numToCount[key3]

			switch {
			case !key3Exists: continue // if third key is not exist, no sense to proceed further
			case !(key1 <= key2 && key2 <= key3): continue// we only need only monotonic triples: key1 <= key2 <= key3
			case key1 == key2 && key1cnt < 2: continue
			case key1 == key3 && key3cnt < 2: continue
			case key2 == key3 && key3cnt < 2: continue
			case key1 == key2 && key2 == key3 && key3cnt < 3: continue
			}

			// if all ok, add unique triple
			ret = append(ret, []int{key1, key2, key3})

		}
	}
	return ret
}

func main() {
	res := threeSum([]int{3, -4, 5, -6, 7, 2, 2, 3, 3, 4, 4, 5, 5, 5})
	fmt.Printf("%v\n", res)
}
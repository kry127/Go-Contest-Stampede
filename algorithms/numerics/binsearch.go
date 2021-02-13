package numerics

/// BINSEARCH ALGORITHM
// author: kry127

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

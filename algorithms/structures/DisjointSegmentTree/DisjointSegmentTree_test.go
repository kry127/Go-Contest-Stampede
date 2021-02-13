package DisjointSegmentTree

import (
	"fmt"
	"math/rand"
	"testing"
)

// segment defined for simple sums
type SumArraySegment struct {
	sum int
}
func (left SumArraySegment) JoinSegments(right Segment) Segment {
	var r SumArraySegment = right.(SumArraySegment)
	return SumArraySegment{left.sum + r.sum}
}

func checkArraySums(array []int) {
	segments := make([]Segment, len(array))
	for i, v := range array {
		segments[i] = SumArraySegment{v}
	}
	dst := BuildDisjointSegmentTree(segments)
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			expected := 0
			for k := i; k <= j; k++ {
				expected += array[k]
			}
			actual := QueryDisjointSegment(dst, i, j).(SumArraySegment).sum
			if expected != actual {
				panic("Sums are not equal")
			}
		}
	}
}


func TestFuzzy(t *testing.T) {
	const MAX_SIZE = 2000
	for n := 0; n < 1000; n++ {
		fmt.Printf("Array %d\n", n + 1)
		size := rand.Uint32() % MAX_SIZE
		arr := make([]int, size)
		for i := range arr {
			arr[i] = rand.Int()
		}
		checkArraySums(arr)
	}
}


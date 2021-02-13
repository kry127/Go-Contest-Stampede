package DisjointSegmentTree


/// DISJOINT SEGMENT TREE GO IMPLEMENTATION
// Author: Kry127

// build complexity: O(n * log n)
// query complexity: O(1)

// To use segment tree, you should produce your own structure like this one:


// define segment interface to operate with disjoint segment tree
type Segment interface {
	JoinSegments(segment2 Segment) Segment
}

// disjoint segment tree -- is a tree of segments
type DisjointSegmentTree = [][]Segment

func highestBit(number int) int {
	height := 0
	for number > 0 {
		height++
		number >>= 1
	}
	return height
}

// calculates polarity and middle for index
func calcMid(lvl, i int) (int, int) {
	polarity := i & (1 << lvl)
	mid := i & (-1 << (lvl + 1)) | (1 << lvl)
	if polarity == 0 {
		return polarity, mid - 1
	} else {
		return polarity, mid
	}
}

// queries segment from the tree
func QueryDisjointSegment(tree DisjointSegmentTree, from, to int) Segment {
	if from > to {
		panic("Invalid query to DisjointSegmentTree")
	}
	if from == to {
		return tree[0][from]
	}

	// main algorithm:
	q := from ^ to
	lvl := highestBit(q) - 1
	row := tree[lvl]
	return row[from].JoinSegments(row[to])
}

func BuildDisjointSegmentTree(input []Segment) DisjointSegmentTree {
	size := len(input)
	height := highestBit(size)

	dst := make([][]Segment, height)
	if size == 0 {
		return dst
	}
	dst[0] = input
	for lvl := 1; lvl < height; lvl++ {
		row := make([]Segment, size)
		for i := 0; i < size; i++ {
			polarity, mid := calcMid(lvl, i)
			// crop mid from overflowing
			if mid >= size {
				mid = size - 1
			}
			if polarity == 0 {
				row[i] = QueryDisjointSegment(dst, i, mid)
			} else {
				row[i] = QueryDisjointSegment(dst, mid, i)
			}
		}
		dst[lvl] = row
	}
	return dst
}
type change struct { start, end, height int }
type heapEntry struct { height, end int }
type maxHeap []heapEntry

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].height > h[j].height }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(heapEntry))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getSkyline(buildings [][]int) [][]int {
    res := [][]int{}
    changes := make([]change, 0, len(buildings) * 2)
    for _, b := range buildings {    
        changes = append(changes, change{b[0], b[1], b[2]}, change{b[1], 0, 0})
    }
    sort.Slice(changes, func(i, j int) bool {
        if changes[i].start == changes[j].start {
            return changes[i].height > changes[j].height
        }
        return changes[i].start < changes[j].start
    })
	res = append(res, []int{0, 0})
	h := &maxHeap{heapEntry{0, math.MaxInt}}
	for _, change := range changes {
        for (*h)[0].end <= change.start {
            heap.Pop(h)
        }
        if change.height > 0 {
            heap.Push(h, heapEntry{change.height, change.end})
        }
        if curHeight := (*h)[0].height; res[len(res) - 1][1] != curHeight {
            res = append(res, []int{change.start, curHeight})
        }
    }
    return res[1:]
}
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
    minHeap := IntHeap{}
	
    for i := 1; i < len(heights); i++ {
        delta := heights[i] - heights[i-1]
        if delta <= 0 {
            continue
        }
        heap.Push(&minHeap, delta)
        if len(minHeap) > ladders {
            bricks -= heap.Pop(&minHeap).(int)
        }
        if bricks < 0 {
            return i - 1
        }
    }
    return len(heights) - 1
}
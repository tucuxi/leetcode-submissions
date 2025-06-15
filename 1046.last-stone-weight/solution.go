func lastStoneWeight(stones []int) int {
    h := make(IntHeap, len(stones))
    copy(h, stones)
    heap.Init(&h)
    for h.Len() > 1 {
        y := heap.Pop(&h).(int)
        x := heap.Pop(&h).(int)
        heap.Push(&h, y - x)
    }
    return h[0]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
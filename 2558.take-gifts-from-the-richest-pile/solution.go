func pickGifts(gifts []int, k int) int64 {
    h := IntHeap(gifts)
    heap.Init(&h)
    for ; k > 0 && len(h) > 0; k-- {
        gift := heap.Pop(&h).(int)
        if rest := int(math.Sqrt(float64(gift))); rest > 0 {
            heap.Push(&h, rest)
        }
    }
    var res int64
    for _, g := range h {
        res += int64(g)
    }
    return res
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
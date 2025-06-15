type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
    idx := make([]int, len(speed))
    for i := range idx {
        idx[i] = i
    }
    sort.Slice(idx, func(i, j int) bool { return efficiency[idx[i]] > efficiency[idx[j]] })
    var pmax int64
    var ssum int64
    h := &IntHeap{}
    for _, v := range idx {
        if h.Len() == k {
            t := heap.Pop(h).(int)
            ssum -= int64(t)
        }
        e, s := efficiency[v], speed[v]
        heap.Push(h, s)
        ssum += int64(s)
        if p := ssum * int64(e); p > pmax {
            pmax = p
        }
    }
    return int(pmax % 1_000_000_007)
}
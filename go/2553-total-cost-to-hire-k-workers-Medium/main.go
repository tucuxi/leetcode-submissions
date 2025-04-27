func totalCost(costs []int, k int, candidates int) int64 {
    var res int64
    p := candidates
    q := len(costs) - min(candidates, len(costs) - candidates)
    first := make(IntHeap, candidates)
    copy(first, costs[:p])
    heap.Init(&first)
    last := make(IntHeap, len(costs) - q)
    copy(last, costs[q:])
    heap.Init(&last)
    for i := 0; i < k; i++ {
        if len(last) == 0 || len(first) > 0 && first[0] <= last[0] {
            res += int64(first[0])
            heap.Pop(&first)
            if p < q {
                heap.Push(&first, costs[p])
                p++
            }
        } else {
            res += int64(last[0])
            heap.Pop(&last)
            if p < q {
                q--
                heap.Push(&last, costs[q])
            }
        }
    }
    return res
}

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

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func minStoneSum(piles []int, k int) int {
    h := make(IntHeap, 0, len(piles))
    heap.Init(&h)
    total := 0
    for _, p := range piles {
        total += p
        heap.Push(&h, p)
    }
    for i := 0; i < k; i++ {
        p := heap.Pop(&h).(int)
        stones := p / 2
        total -= stones
        heap.Push(&h, p - stones)
    }
    return total
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
	x := old[len(old) - 1]
	*h = old[0 : len(old) - 1]
	return x
}
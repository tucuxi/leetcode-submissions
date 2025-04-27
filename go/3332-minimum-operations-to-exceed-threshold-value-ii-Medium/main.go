func minOperations(nums []int, k int) int {
    var h IntHeap = nums
    heap.Init(&h)
    i := 0
    for ; h[0] < k; i++ {
        a := heap.Pop(&h).(int)
        b := heap.Pop(&h).(int)
        heap.Push(&h, a * 2 + b)
    }
    return i
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
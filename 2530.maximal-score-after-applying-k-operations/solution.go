func maxKelements(nums []int, k int) int64 {
    maxHeap := &IntHeap{}
    for _, num := range nums {
        heap.Push(maxHeap, num)
    }
    var score int64
    for i := 0; i < k; i++ {
        num := heap.Pop(maxHeap).(int)
        score += int64(num)
        heap.Push(maxHeap, (num + 2) /3)
    }
    return score
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
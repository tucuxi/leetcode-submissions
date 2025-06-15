func maximumScore(a int, b int, c int) int {
    res := 0
    h := &maxHeap{a, b, c}
    heap.Init(h)
    for {
        stone1 := heap.Pop(h).(int)
        stone2 := heap.Pop(h).(int)
        if stone2 == 0 {
            return res
        }
        res++
        heap.Push(h, stone1 - 1)
        heap.Push(h, stone2 - 1)
    }
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
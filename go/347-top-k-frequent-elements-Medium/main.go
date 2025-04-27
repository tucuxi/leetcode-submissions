func topKFrequent(nums []int, k int) []int {
    frequencies := map[int]int{}
    for _, n := range nums {
        frequencies[n]++
    }
    h := &MaxHeap{}
    for n, c := range frequencies {
        heap.Push(h, [2]int{n, c})
    }
    res := make([]int, 0, k)
    for i := 0; i < k; i++ {
        element := heap.Pop(h).([2]int)
        res = append(res, element[0])
    }
    return res
}

type MaxHeap [][2]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
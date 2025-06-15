type maxHeap [][2]int
type any=interface{}

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
    *h = append(*h, x.([2]int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxResult(nums []int, k int) int {
    h := &maxHeap{ [2]int{0, nums[0]} }
    max := nums[0]

    for i := 1; i < len(nums); i++ {
        t := (*h)[0]
        for t[0] < i-k {
            heap.Pop(h)
            t = (*h)[0]
        }
        max = t[1] + nums[i]
        heap.Push(h, [2]int{i, max})
    }
    return max
}
func shortestSubarray(nums []int, k int) int {
    res := math.MaxInt
    h := MinHeap{}
    sum := 0
    for i := range nums {
        sum += nums[i]
        if sum >= k {
            res = min(res, i + 1)
        }
        for len(h) > 0 && sum - h[0][0] >= k {
            res = min(res, i - h[0][1])
            heap.Pop(&h)
        }
        heap.Push(&h, [2]int{sum, i})
    }
    if res == math.MaxInt {
        return -1
    }
    return res
}

type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
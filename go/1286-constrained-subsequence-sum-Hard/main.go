func constrainedSubsetSum(nums []int, k int) int {
    h := make(maxHeap, 0, k)
    heap.Push(&h, sumWithIndex{nums[0], 0})
    res := nums[0]
    for i := 1; i < len(nums); i++ {
        for i - h[0].index > k {
            heap.Pop(&h)
        }
        cur := nums[i]
        if h[0].sum > 0 {
            cur += h[0].sum
        }
        if cur > res {
            res = cur
        }
        heap.Push(&h, sumWithIndex{cur, i})
    }
    return res
}

type sumWithIndex struct{sum, index int}
type maxHeap []sumWithIndex

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].sum > h[j].sum }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(sumWithIndex))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

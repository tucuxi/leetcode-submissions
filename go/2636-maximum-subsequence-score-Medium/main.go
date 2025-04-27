func maxScore(nums1 []int, nums2 []int, k int) int64 {
    idx := make([]int, len(nums1))
    for i := range idx {
        idx[i] = i
    }
    sort.Slice(idx, func(i, j int) bool { return nums2[idx[i]] > nums2[idx[j]] })
    nums1Heap := make(minHeap, 0, len(nums1))
    sum := 0
    for i := 0; i < k; i++ {
        n := nums1[idx[i]]
        sum += n
        heap.Push(&nums1Heap, n)
    }
    res := int64(sum) * int64(nums2[idx[k - 1]])
    for i := k; i < len(nums1); i++ {
        smallest := heap.Pop(&nums1Heap).(int)
        n := nums1[idx[i]]
        heap.Push(&nums1Heap, n)
        sum += n - smallest 
        if cand := int64(sum) * int64(nums2[idx[i]]); cand > res {
            res = cand
        }
    }
    return res                                 
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

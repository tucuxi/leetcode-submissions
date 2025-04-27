func getFinalState(nums []int, k int, multiplier int) []int {
    h := make(indexedValueHeap, len(nums))
    for i := range nums {
        h[i] = indexedValue{i, nums[i]}
    }
    heap.Init(&h)
    for range k {
        minElem := heap.Pop(&h).(indexedValue)
        minElem.value *= multiplier
        nums[minElem.index] = minElem.value
        heap.Push(&h, minElem)
    }
    return nums    
}

type indexedValue struct { index, value int }
type indexedValueHeap []indexedValue

func (h indexedValueHeap) Len() int           { return len(h) }
func (h indexedValueHeap) Less(i, j int) bool { return h[i].value < h[j].value || h[i].value == h[j].value && h[i].index < h[j].index }
func (h indexedValueHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *indexedValueHeap) Push(x any) {
	*h = append(*h, x.(indexedValue))
}

func (h *indexedValueHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
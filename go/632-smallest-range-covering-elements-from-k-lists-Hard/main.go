func smallestRange(nums [][]int) []int {
    h := make(IndexedIntHeap, len(nums))
    maxValue := math.MinInt
    for i, numList := range nums {
        h[i] = [2]int{i, numList[0]}
        maxValue = max(numList[0], maxValue)
    }
    heap.Init(&h)
    
    indices := make([]int, len(nums))    
    a, b := 0, math.MaxInt

    for {
        minElement := heap.Pop(&h).([2]int)
        list, minValue := minElement[0], minElement[1]
        if maxValue - minValue < b - a {
            a, b = minValue, maxValue
        }
        indices[list]++
        if indices[list] == len(nums[list]) {
            return []int{a, b}
        }
        nextInList := nums[list][indices[list]]
        heap.Push(&h, [2]int{list, nextInList})
        maxValue = max(nextInList, maxValue)
    }
}

type IndexedIntHeap [][2]int

func (h IndexedIntHeap) Len() int           { return len(h) }
func (h IndexedIntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IndexedIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IndexedIntHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *IndexedIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
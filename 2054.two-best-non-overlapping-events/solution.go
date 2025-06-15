func maxTwoEvents(events [][]int) int {
    slices.SortFunc(events, func(a, b []int) int { return a[0] - b[0] })
    
    var h entryHeap
    maxValue := 0
    maxSum := 0

    for _, event := range events {
        for len(h) > 0 && h[0].endTime < event[0] {
            h0 := heap.Pop(&h).(entry)
            maxValue = max(maxValue, h0.value)
        }
        maxSum = max(maxSum, maxValue + event[2])
        heap.Push(&h, entry{event[1], event[2]})
    }
    
    return maxSum
}

type entry struct {
    endTime int
    value int
}

type entryHeap []entry

func (h entryHeap) Len() int           { return len(h) }
func (h entryHeap) Less(i, j int) bool { return h[i].endTime < h[j].endTime }
func (h entryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *entryHeap) Push(x any) {
	*h = append(*h, x.(entry))
}

func (h *entryHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
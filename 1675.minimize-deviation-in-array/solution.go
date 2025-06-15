func minimumDeviation(nums []int) int {
    var h maxHeap
    min := math.MaxInt
    for _, n := range nums {
        if n & 1 != 0 {
            n *= 2
        }
        heap.Push(&h, n)
        if n < min {
            min = n
        }
    }
    res := math.MaxInt
    for {
        max := heap.Pop(&h).(int)
        if d := max - min; d < res {
            res = d
        }
        if max & 1 != 0 {
            break
        }
        max /= 2
        heap.Push(&h, max)
        if max < min {
            min = max
        }
    }
    return res
}

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {	*h = append(*h, x.(int)) }

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
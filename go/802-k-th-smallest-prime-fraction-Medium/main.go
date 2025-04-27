func kthSmallestPrimeFraction(arr []int, k int) []int {
    lastIndex := len(arr) - 1
    var h fractionHeap

    for i := range lastIndex {
        heap.Push(&h, fraction{arr[i], arr[lastIndex], lastIndex})
    }
    for ; k > 1; k-- {
        f := heap.Pop(&h).(fraction)
        f.denominatorIndex--
        f.denominator = arr[f.denominatorIndex]
        if f.nominator < f.denominator {
            heap.Push(&h, f)
        }
    }

    f := heap.Pop(&h).(fraction)
    return []int{f.nominator, f.denominator}
}

type fraction struct {
    nominator int
    denominator int
    denominatorIndex int
}

type fractionHeap []fraction

func (h fractionHeap) Len() int           { return len(h) }
func (h fractionHeap) Less(i, j int) bool { return h[i].nominator * h[j].denominator < h[j].nominator * h[i].denominator }
func (h fractionHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *fractionHeap) Push(x any) {
	*h = append(*h, x.(fraction))
}

func (h *fractionHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
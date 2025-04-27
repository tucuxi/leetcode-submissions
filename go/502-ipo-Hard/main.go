func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
    var ph maxProfitHeap
    var ch minCapitalHeap

    for i := range profits {
        p := project{profits[i], capital[i]}
        if capital[i] <= w {
            heap.Push(&ph, p)
        } else {
            heap.Push(&ch, p)
        }
    }
    for i := 0; i < k && len(ph) > 0; i++ {
        p := heap.Pop(&ph).(project)
        w += p.profit
        for len(ch) > 0 && ch[0].capital <= w {
            heap.Push(&ph, heap.Pop(&ch))
        }
    }
    return w
}

type project struct{profit, capital int}
type maxProfitHeap []project
type minCapitalHeap []project

func (h maxProfitHeap) Len() int           { return len(h) }
func (h maxProfitHeap) Less(i, j int) bool { return h[i].profit > h[j].profit }
func (h maxProfitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxProfitHeap) Push(x interface{}) {
	*h = append(*h, x.(project))
}

func (h *maxProfitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h minCapitalHeap) Len() int           { return len(h) }
func (h minCapitalHeap) Less(i, j int) bool { return h[i].capital <  h[j].capital }
func (h minCapitalHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minCapitalHeap) Push(x interface{}) {
	*h = append(*h, x.(project))
}

func (h *minCapitalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
    h := make(maxGainHeap, len(classes))
    for i, c := range classes {
        h[i] = entry{pass: c[0], total: c[1], gain: gain(c[0], c[1])}
    }
    heap.Init(&h)
    for range extraStudents {
        e := heap.Pop(&h).(entry)
        heap.Push(&h, entry{pass: e.pass + 1, total: e.total + 1, gain: gain(e.pass + 1, e.total + 1)})
    }
    var passRatioSum float64
    for h.Len() > 0 {
        e := heap.Pop(&h).(entry)
        passRatioSum += float64(e.pass) / float64(e.total)
    }
    return passRatioSum / float64(len(classes))
}

func gain(pass, total int) float64 {
    return float64(pass + 1) / float64(total + 1) - float64(pass) / float64(total) 
}

type entry struct {
    pass, total int
    gain float64
}

type maxGainHeap []entry

func (h maxGainHeap) Len() int           { return len(h) }
func (h maxGainHeap) Less(i, j int) bool { return h[i].gain > h[j].gain }
func (h maxGainHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxGainHeap) Push(x any) {
	*h = append(*h, x.(entry))
}

func (h *maxGainHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
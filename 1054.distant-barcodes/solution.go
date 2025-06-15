func rearrangeBarcodes(barcodes []int) []int {
    h := func() IntHeap {
        f := make(map[int]int)
        for _, code := range barcodes {
            f[code]++
        }
        h := make(IntHeap, 0, len(f))
        for code, freq := range f {
            h = append(h, occurrence{code, freq})
        }
        heap.Init(&h)
        return h
    }()
    
    var prev occurrence
    
    for i := range barcodes {
        curr := heap.Pop(&h).(occurrence)
        barcodes[i] = curr.code
        curr.freq--
        if prev.freq > 0 {
            heap.Push(&h, prev)
        }
        prev = curr
    }
    return barcodes
}

type occurrence struct { code, freq int }

type IntHeap []occurrence

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(occurrence))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
type KthLargest struct {
    largestK IntHeap
}


func Constructor(k int, nums []int) KthLargest {
    kth := KthLargest{make(IntHeap, 0, k)}
    for _, n := range nums {
        kth.Add(n)
    }
    return kth
}


func (this *KthLargest) Add(val int) int {
    if len(this.largestK) < cap(this.largestK) {
        heap.Push(&this.largestK, val)
    } else if val > this.largestK[0] {
        heap.Pop(&this.largestK)        
        heap.Push(&this.largestK, val)
    }
    return this.largestK[0]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SmallestInfiniteSet struct {
    next int
    addedBackHeap IntHeap
    addedBackSet map[int]bool
}


func Constructor() SmallestInfiniteSet {
    return SmallestInfiniteSet{
        next: 1,
        addedBackHeap: make(IntHeap, 0, 100),
        addedBackSet: make(map[int]bool),
    }
}


func (this *SmallestInfiniteSet) PopSmallest() int {
    if this.addedBackHeap.Len() > 0 && this.addedBackHeap[0] < this.next {
        res := heap.Pop(&this.addedBackHeap).(int)
        delete(this.addedBackSet, res)
        return res
    }
    res := this.next
    this.next++
    return res
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    if num < this.next && !this.addedBackSet[num] {
        this.addedBackSet[num] = true
        heap.Push(&this.addedBackHeap, num)
    }
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
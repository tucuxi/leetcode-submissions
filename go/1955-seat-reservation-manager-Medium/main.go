type SeatManager struct {
    returnedSeats IntHeap
    nextFreeSeat int
}

func Constructor(n int) SeatManager {
    return SeatManager{nextFreeSeat: 1}
}

func (this *SeatManager) Reserve() int {
    if len(this.returnedSeats) > 0 {
        return heap.Pop(&this.returnedSeats).(int)
    }
    seatNumber := this.nextFreeSeat
    this.nextFreeSeat++
    return seatNumber    
}

func (this *SeatManager) Unreserve(seatNumber int)  {
    heap.Push(&this.returnedSeats, seatNumber)
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
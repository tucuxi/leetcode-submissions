type event struct {
    time int
    friend int
    arriving bool
}

func smallestChair(times [][]int, targetFriend int) int {
    var availableChairs IntHeap
    maxChair := -1

    events := make([]event, 0, 2 * len(times))
    for i, t := range times {
        events = append(events, event{t[0], i, true}, event{t[1], i, false})
    }
    slices.SortFunc(events, func(a, b event) int {
        if a.time == b.time && b.arriving {
            return -1
        }
        return a.time - b.time
    })

    assignedChairs := make([]int, len(times))

    for _, e := range events {
        if e.arriving {
            var chair int
            if len(availableChairs) == 0 {
                maxChair++
                chair = maxChair
            } else {
                chair = heap.Pop(&availableChairs).(int)
            }
            assignedChairs[e.friend] = chair
        } else {
            chair := assignedChairs[e.friend]
            if chair == maxChair {
                maxChair--
            } else {
                heap.Push(&availableChairs, chair)
            }
        }
    }
    return assignedChairs[targetFriend]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
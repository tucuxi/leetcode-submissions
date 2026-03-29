type Event struct {
    id       int
    priority int
    index    int
}

type EventPriorityQueue []*Event

func (pq EventPriorityQueue) Len() int { return len(pq) }

func (pq EventPriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority || pq[i].priority == pq[j].priority && pq[i].id < pq[j].id
}

func (pq EventPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *EventPriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Event)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *EventPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	event := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	event.index = -1 // for safety
	*pq = old[0 : n-1]
	return event
}

type EventManager struct {
    mp map[int]*Event
    pq EventPriorityQueue
}


func Constructor(events [][]int) EventManager {
    mp := make(map[int]*Event)
    pq := make(EventPriorityQueue, len(events))
    for i := range events {
        id := events[i][0]
        event := &Event{id, events[i][1], i}
        mp[id] = event
        pq[i] = event
    }
    heap.Init(&pq)
    return EventManager{mp, pq}
}


func (this *EventManager) UpdatePriority(eventId int, newPriority int)  {
    event := this.mp[eventId]
    event.priority = newPriority
    heap.Fix(&this.pq, event.index)
}


func (this *EventManager) PollHighest() int {
    if this.pq.Len() == 0 {
        return -1
    }
    event := heap.Pop(&this.pq).(*Event)
    delete(this.mp, event.id)
    return event.id
}


/**
 * Your EventManager object will be instantiated and called as such:
 * obj := Constructor(events);
 * obj.UpdatePriority(eventId,newPriority);
 * param_2 := obj.PollHighest();
 */
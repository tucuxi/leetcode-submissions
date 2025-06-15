type item struct {
    key int
    value int
    count int
    time int64
    index int
}

type priorityQueue []*item

type LFUCache struct {
    capacity int
    entries map[int]*item
    queue priorityQueue
    time int64
}

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].count < pq[j].count || pq[i].count == pq[j].count && pq[i].time < pq[j].time
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        capacity: capacity,
        entries: make(map[int]*item, capacity),
        queue: make(priorityQueue, 0, capacity),
        time: 0,
    }    
}

func (this *LFUCache) Get(key int) int {
    if entry, found := this.entries[key]; found {
        this.time++
        entry.count++
        entry.time = this.time
        heap.Fix(&this.queue, entry.index)
        return entry.value
    } 
    return -1
}

func (this *LFUCache) Put(key int, value int) {
    if this.capacity == 0 {
        return
    }
    this.time++
    if entry, found := this.entries[key]; found {
        entry.value = value
        entry.count++
        entry.time = this.time
        heap.Fix(&this.queue, entry.index)
        return
    }
    if len(this.queue) == this.capacity {
        victim := heap.Pop(&this.queue).(*item)
        delete(this.entries, victim.key)
    }
    entry := &item{
        key: key,
        value: value,
        count: 1,
        time: this.time,
        index: -1,
    }
    this.entries[key] = entry
    heap.Push(&this.queue, entry)
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
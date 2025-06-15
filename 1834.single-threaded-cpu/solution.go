func getOrder(tasks [][]int) []int {
    res := make([]int, 0, len(tasks))
    scheduled, waiting := make(ScheduledQueue, len(tasks)), make(WaitingQueue, 0)
    for i := range tasks {
       scheduled[i] = &task{i, tasks[i][0], tasks[i][1], -1}
    }
    heap.Init(&scheduled)
    time := 0
    for len(res) < len(tasks) {
        for scheduled.Len() > 0 {
            next := heap.Pop(&scheduled).(*task)
            if next.enqueueTime <= time  {
                heap.Push(&waiting, next)
            } else {
                heap.Push(&scheduled, next)
                break
            }
        }
        if waiting.Len() == 0 {
            next := heap.Pop(&scheduled).(*task)
            heap.Push(&waiting, next)
            time = next.enqueueTime
            continue
        }
        next := heap.Pop(&waiting).(*task)
        res = append(res, next.taskId)
        time += next.processingTime
    }
    return res
}

type task struct {
    taskId int
	enqueueTime int
    processingTime int
	index int
}

type ScheduledQueue []*task

func (pq ScheduledQueue) Len() int { return len(pq) }

func (pq ScheduledQueue) Less(i, j int) bool {
	return pq[i].enqueueTime < pq[j].enqueueTime 
}

func (pq ScheduledQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *ScheduledQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*task)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *ScheduledQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type WaitingQueue []*task

func (pq WaitingQueue) Len() int { return len(pq) }

func (pq WaitingQueue) Less(i, j int) bool {
	return pq[i].processingTime < pq[j].processingTime ||
        pq[i].processingTime == pq[j].processingTime && pq[i].taskId < pq[j].taskId
}

func (pq WaitingQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *WaitingQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*task)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *WaitingQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
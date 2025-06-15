type TaskManager struct {
    tasksById map[int]*task
    priorityQueue PriorityQueue
}


func Constructor(tasks [][]int) TaskManager {
    tasksById := make(map[int]*task)
    priorityQueue := make(PriorityQueue, len(tasks))
    for i, t := range tasks {
        item := &task{taskId: t[1], userId: t[0], priority: t[2], deleted: false, index: i}
        tasksById[item.taskId] = item
        priorityQueue[i] = item
    }
    heap.Init(&priorityQueue)
    return TaskManager{tasksById, priorityQueue}
}


func (this *TaskManager) Add(userId int, taskId int, priority int)  {
    item := &task{taskId: taskId, userId: userId, priority: priority}
    this.tasksById[taskId] = item
    heap.Push(&this.priorityQueue, item)
}


func (this *TaskManager) Edit(taskId int, newPriority int)  {
    item := this.tasksById[taskId]
	item.priority = newPriority
	heap.Fix(&this.priorityQueue, item.index)
}


func (this *TaskManager) Rmv(taskId int)  {
    item := this.tasksById[taskId]
    item.deleted = true
    delete(this.tasksById, taskId)
}


func (this *TaskManager) ExecTop() int {
    for len(this.priorityQueue) > 0 {
        item := heap.Pop(&this.priorityQueue).(*task)
        if !item.deleted {
            delete(this.tasksById, item.taskId)
            return item.userId
        }
    }
    return -1
}


/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */

type task struct {
    taskId, userId, priority, index int
    deleted bool
}

type PriorityQueue []*task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority || pq[i].priority == pq[j].priority && pq[i].taskId > pq[j].taskId
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*task)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
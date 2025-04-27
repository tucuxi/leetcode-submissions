type MyCircularQueue struct {
    items []int
    used int
    next int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{items: make([]int, k), used: 0, next: 0}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.used == cap(this.items) {
        return false
    }
    this.items[this.next] = value
    this.used++
    this.next = (this.next + 1) % cap(this.items)
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.used == 0 {
        return false
    }
    this.used--
    return true
}

func (this *MyCircularQueue) Front() int {
    if this.used == 0 {
        return -1
    }
    i := this.next - this.used
    if i < 0 {
        i += cap(this.items)
    }
    return this.items[i]
}

func (this *MyCircularQueue) Rear() int {
    if this.used == 0 {
        return -1
    }
    i := this.next - 1
    if i < 0 {
        i += cap(this.items)
    }
    return this.items[i]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.used == 0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.used == cap(this.items)
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
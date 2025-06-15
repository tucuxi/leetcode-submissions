type MyCircularDeque struct {
    elements []int
    size int
    p, q int
}


func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{elements: make([]int, k), size: 0, p: k - 1, q: 0}
}


func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.size == len(this.elements) {
        return false
    }
    this.elements[this.p] = value
    this.p = modDecrement(this.p, len(this.elements))
    this.size++
    return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.size == len(this.elements) {
        return false
    }
    this.elements[this.q] = value
    this.q = modIncrement(this.q, len(this.elements))
    this.size++
    return true
}


func (this *MyCircularDeque) DeleteFront() bool {
    if this.size == 0 {
        return false
    }
    this.p = modIncrement(this.p, len(this.elements))
    this.size--
    return true
}


func (this *MyCircularDeque) DeleteLast() bool {
    if this.size == 0 {
        return false
    }
    this.q = modDecrement(this.q, len(this.elements))
    this.size--
    return true
}


func (this *MyCircularDeque) GetFront() int {
    if this.size == 0 {
        return -1
    }
    return this.elements[modIncrement(this.p, len(this.elements))]
}


func (this *MyCircularDeque) GetRear() int {
    if this.size == 0 {
        return -1
    }
    return this.elements[modDecrement(this.q, len(this.elements))]
}


func (this *MyCircularDeque) IsEmpty() bool {
    return this.size == 0
}


func (this *MyCircularDeque) IsFull() bool {
    return this.size == len(this.elements) 
}


func modDecrement(value, k int) int {
    if value == 0 {
        return k - 1
    }
    return value - 1
}

func modIncrement(value, k int) int {
    if value == k - 1 {
        return 0
    }
    return value + 1
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
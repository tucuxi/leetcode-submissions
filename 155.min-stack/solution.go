type MinStack struct {
    head *entry
}

type entry struct {
    val  int
    min  int
    next *entry
}

func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    min := val
    if this.head != nil && this.head.min < min {
        min = this.head.min
    }
    this.head = &entry{val, min, this.head}
}


func (this *MinStack) Pop()  {
    this.head = this.head.next
}


func (this *MinStack) Top() int {
    return this.head.val
}


func (this *MinStack) GetMin() int {
    return this.head.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
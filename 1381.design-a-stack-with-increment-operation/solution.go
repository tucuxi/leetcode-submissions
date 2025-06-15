type CustomStack struct {
    elements []int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{elements: make([]int, 0, maxSize)}
}


func (this *CustomStack) Push(x int)  {
    e := this.elements
    if len(e) < cap(e) {
        this.elements = append(e, x)
    }
}


func (this *CustomStack) Pop() int {
    e := this.elements
    if len(e) == 0 {
        return -1
    }
    top := e[len(e) - 1]
    this.elements = e[:len(e) - 1]
    return top
}


func (this *CustomStack) Increment(k int, val int)  {
    e := this.elements
    for i := range min(k, len(e)) {
        e[i] += val
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
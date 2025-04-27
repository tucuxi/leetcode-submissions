type MyStack struct {
    elems []int
}

func Constructor() MyStack {
    return MyStack{}
}

func (this *MyStack) Push(x int)  {
    this.elems = append(this.elems, x)
    for i := len(this.elems); i > 1; i-- {
        this.elems = append(this.elems, this.elems[0])
        this.elems = this.elems[1:]
    }
}

func (this *MyStack) Pop() int {
    res := this.elems[0]
    this.elems = this.elems[1:]
    return res
}

func (this *MyStack) Top() int {
    return this.elems[0]
}

func (this *MyStack) Empty() bool {
    return len(this.elems) == 0
}
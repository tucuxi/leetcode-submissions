type MyQueue struct {
    in, out stack
}


func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    this.in.push(x)
}


func (this *MyQueue) Pop() int {
    this.shuffle()
    return this.out.pop()
}

func (this *MyQueue) Peek() int {
    this.shuffle()
    return this.out.peek() 
}

func (this *MyQueue) Empty() bool {
    this.shuffle()
    return this.out.empty()
}

func (this *MyQueue) shuffle() {
    if this.out.empty() {
        for !this.in.empty() {
            this.out.push(this.in.pop())
        }
    }
}

type stack []int

func (this *stack) push(x int) {
    *this = append(*this, x)
}

func (this *stack) pop() int {
    res := (*this)[len(*this) - 1]
    *this = (*this)[:len(*this) - 1]
    return res
}

func (this *stack) peek() int {
    return (*this)[len(*this) - 1]
}

func (this *stack) empty() bool {
    return len(*this) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
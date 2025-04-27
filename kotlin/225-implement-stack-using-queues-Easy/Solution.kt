class MyStack() {

    private val q = ArrayDeque<Int>()
    
    fun push(x: Int) {
        q.offer(x)
        repeat(q.size - 1) { 
            q.offer(q.poll())
        }
    }

    fun pop(): Int = q.poll()

    fun top(): Int = q.peek()

    fun empty(): Boolean = q.isEmpty()

}

/**
 * Your MyStack object will be instantiated and called as such:
 * var obj = MyStack()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.empty()
 */
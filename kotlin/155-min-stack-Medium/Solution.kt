class MinStack() {

    private class Node(val x: Int, val minx: Int, val next: Node?)
    
    private var head: Node? = null
    
    fun push(x: Int) {
        val prevMin = head?.minx ?: x
        head = Node(x, minOf(x, prevMin), head)
    }

    fun pop() {
        head = head!!.next
    }

    fun top(): Int {
        return head!!.x
    }

    fun getMin(): Int {
        return head!!.minx
    }

}

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = MinStack()
 * obj.push(`val`)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */
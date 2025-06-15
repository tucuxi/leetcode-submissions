import scala.collection.mutable.Stack

class MyQueue() {

    /** Initialize your data structure here. */
    private val sinp = Stack[Int]()
    private val sout = Stack[Int]()
    
    /** Push element x to the back of queue. */
    def push(x: Int): Unit = {
        sinp.push(x)
    }

    /** Removes the element from in front of queue and returns that element. */
    def pop(): Int = {
        shuffle
        sout.pop
    }

    /** Get the front element. */
    def peek(): Int = {
        shuffle
        sout.top
    }

    /** Returns whether the queue is empty. */
    def empty(): Boolean = {
        shuffle
        sout.isEmpty
    }
        
    private def shuffle(): Unit = {
        if (sout.isEmpty)
            while (sinp.nonEmpty)
                sout.push(sinp.pop)
    }
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * var obj = new MyQueue()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.peek()
 * var param_4 = obj.empty()
 */
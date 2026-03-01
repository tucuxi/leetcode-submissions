class FrontMiddleBackQueue() {

    private val front = ArrayDeque<Int>()
    private val back = ArrayDeque<Int>()

    fun pushFront(`val`: Int) {
        front.addFirst(`val`)
        balance()
    }

    fun pushMiddle(`val`: Int) {
        back.addFirst(`val`)
        balance()        
    }

    fun pushBack(`val`: Int) {
        back.addLast(`val`)
        balance()
    }

    fun popFront() = (front.removeFirstOrNull() ?: back.removeFirstOrNull() ?: -1).also { balance() }

    fun popMiddle(): Int {
        val res = if (front.size == back.size) front.removeLastOrNull() else back.removeFirstOrNull()
        balance()
        return res ?: -1
    }

    fun popBack() = back.removeLastOrNull()?.also { balance() } ?: -1

    private fun balance() {
        when(front.size - back.size) {
            -2 -> front.addLast(back.removeFirst())
             1 -> back.addFirst(front.removeLast())
        }
    }
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * var obj = FrontMiddleBackQueue()
 * obj.pushFront(`val`)
 * obj.pushMiddle(`val`)
 * obj.pushBack(`val`)
 * var param_4 = obj.popFront()
 * var param_5 = obj.popMiddle()
 * var param_6 = obj.popBack()
 */
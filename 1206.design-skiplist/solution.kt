const val MAX_LEVEL = 15

class Skiplist() {

    private val sentinel = SkipNode(0, MAX_LEVEL)

    private fun findPredecessors(target: Int): Array<SkipNode<Int>> {
        val preds = Array(MAX_LEVEL + 1) { sentinel }
        var node = sentinel
        
        for (i in MAX_LEVEL downTo 0) {
            var nextNode = node.next[i]

            while (nextNode != null && nextNode.value < target) {
                node = nextNode
                nextNode = node.next[i]
            }
            preds[i] = node
        }
        return preds
    }

    fun search(target: Int): Boolean {
        var node = sentinel
        
        for (i in MAX_LEVEL downTo 0) {
            var nextNode = node.next[i]

            while (nextNode != null && nextNode.value < target) {
                node = nextNode
                nextNode = node.next[i]
            }
        }
        val candidate = node.next[0]
        return candidate != null && candidate.value == target
    }

    fun add(num: Int) {
        val preds = findPredecessors(num)
        val newLevel = randomLevel()
        val newNode = SkipNode(num, newLevel)

        for (i in 0..newLevel) {
            newNode.next[i] = preds[i].next[i]
            preds[i].next[i] = newNode
        }
    }

    fun erase(num: Int): Boolean {
        val preds = findPredecessors(num)
        val candidate = preds[0].next[0]

        if (candidate == null || candidate.value != num) {
            return false
        }
        for (i in 0..MAX_LEVEL) {
            if (preds[i].next[i] != candidate) {
                break
            }
            preds[i].next[i] = candidate.next[i]
        }
        return true
    }

    private fun randomLevel(): Int {
        var level = 0
        while (level < MAX_LEVEL && Random.nextBoolean()) {
            level++
        }
        return level
    }
}

class SkipNode<T : Comparable<T>>(val value: T, val level: Int) {
    val next = arrayOfNulls<SkipNode<T>>(level + 1)
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * var obj = Skiplist()
 * var param_1 = obj.search(target)
 * obj.add(num)
 * var param_3 = obj.erase(num)
 */
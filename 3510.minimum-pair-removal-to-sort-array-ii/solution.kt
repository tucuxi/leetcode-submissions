class Solution {

    data class Node(
        var value: Long,
        val originalIndex: Int,
        var prev: Node? = null,
        var next: Node? = null,
        var version: Int = 0,
        var removed: Boolean = false
    )

    data class PairInfo(
        val sum: Long,
        val nodeIndex: Int,
        val node: Node,
        val lVersion: Int,
        val rVersion: Int
    ) : Comparable<PairInfo> {
        override fun compareTo(other: PairInfo): Int {
            val sumCmp = this.sum.compareTo(other.sum)
            if (sumCmp != 0) return sumCmp
            return this.nodeIndex.compareTo(other.nodeIndex)
        }
    }

    fun minimumPairRemoval(nums: IntArray): Int {
        if (nums.isEmpty()) return 0

        val nodes = nums.mapIndexed { index, value -> Node(value.toLong(), index) }
        for (i in 0 until nodes.lastIndex) {
            nodes[i].next = nodes[i + 1]
            nodes[i + 1].prev = nodes[i]
        }

        var descents = 0
        var curr = nodes[0]
        while (curr.next != null) {
            if (curr.value > curr.next!!.value) {
                descents++
            }
            curr = curr.next!!
        }

        if (descents == 0) return 0

        val pq = PriorityQueue<PairInfo>()
        curr = nodes[0]
        while (curr.next != null) {
            val nextNode = curr.next!!
            pq.add(PairInfo(
                sum = curr.value + nextNode.value,
                nodeIndex = curr.originalIndex,
                node = curr,
                lVersion = curr.version,
                rVersion = nextNode.version
            ))
            curr = nextNode
        }

        var operations = 0

        while (descents > 0 && !pq.isEmpty()) {
            val minPair = pq.poll()
            val L = minPair.node
            val R = L.next

            if (L.removed || R == null || R.removed) continue
            if (L.version != minPair.lVersion || R.version != minPair.rVersion) continue

            operations++

            if (L.prev != null && L.prev!!.value > L.value) descents--
            if (L.value > R.value) descents--
            if (R.next != null && R.value > R.next!!.value) descents--

            val newVal = L.value + R.value
            L.value = newVal
            L.version++

            R.removed = true
            L.next = R.next
            if (L.next != null) {
                L.next!!.prev = L
            }

            if (L.prev != null && L.prev!!.value > L.value) descents++
            if (L.next != null && L.value > L.next!!.value) descents++

            if (L.prev != null) {
                pq.add(PairInfo(
                    sum = L.prev!!.value + L.value,
                    nodeIndex = L.prev!!.originalIndex,
                    node = L.prev!!,
                    lVersion = L.prev!!.version,
                    rVersion = L.version
                ))
            }
            if (L.next != null) {
                pq.add(PairInfo(
                    sum = L.value + L.next!!.value,
                    nodeIndex = L.originalIndex,
                    node = L,
                    lVersion = L.version,
                    rVersion = L.next!!.version
                ))
            }
        }

        return operations
    }
}
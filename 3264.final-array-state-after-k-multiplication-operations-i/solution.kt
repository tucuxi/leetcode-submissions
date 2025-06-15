class Solution {
    fun getFinalState(nums: IntArray, k: Int, multiplier: Int): IntArray {
        val pq = PriorityQueue<Pair<Int, Int>>(compareBy({ it.second }, { it.first }))
        nums.forEachIndexed { index, value -> pq.offer(index to value) }
        repeat(k) {
            val (minIndex, minValue) = pq.poll()
            pq.offer(minIndex to minValue * multiplier) 
        }
        val res = IntArray(nums.size)
        while (pq.isNotEmpty()) {
            val (index, value) = pq.poll()
            res[index] = value
        }
        return res
    }
}
class Solution {
    fun maxKelements(nums: IntArray, k: Int): Long {
        val pq = PriorityQueue<Int>(Comparator.reverseOrder())
        pq.addAll(nums.toTypedArray())
        var score = 0L
        repeat(k) {
            val n = pq.poll()
            score += n
            pq.offer((n + 2) / 3)
        }
        return score
    }
}
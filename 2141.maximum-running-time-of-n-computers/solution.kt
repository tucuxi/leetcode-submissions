class Solution {
    fun maxRunTime(n: Int, batteries: IntArray): Long {
        val power = batteries.sumOf { it.toLong() }
        var lo = 1L
        var hi = power / n
        while (lo < hi) {
            val target = hi - (hi - lo) / 2
            val extra = batteries.sumOf { minOf(it.toLong(), target) }
            if (extra >= n * target) {
                lo = target
            } else {
                hi = target - 1
            }
        }
        return lo
    }
}
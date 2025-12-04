const val MOD = 1_000_000_007

class Solution {
    fun countTrapezoids(points: Array<IntArray>): Int {
        val levelCount = points.groupingBy { it[1] }.eachCount()
        val (_, res) = levelCount.values.fold(0L to 0L) { (sum, res), count ->
            val pairs = count.toLong() * (count - 1) / 2
            (sum + pairs) % MOD to (res + pairs * sum) % MOD
        }
        return res.toInt()
    }
}
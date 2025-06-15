class Solution {
    fun distributeCandies(n: Int, limit: Int): Long {
        var res = 0L
        for (i in 0..minOf(limit, n)) {
            val remaining = n - i
            if (remaining <= 2 * limit) {
                res += (minOf(remaining, limit) - maxOf(0, remaining - limit) + 1).toLong()
            }
        }
        return res
    }
}
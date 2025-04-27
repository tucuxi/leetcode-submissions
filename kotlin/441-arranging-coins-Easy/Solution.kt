class Solution {
    fun arrangeCoins(n: Int): Int {
        var lo = 1L
        var hi = n.toLong()
        while (lo <= hi) {
            val m = lo + (hi - lo) / 2
            val coins = m * (m + 1) / 2
            if (coins <= n) {
                lo = m + 1
            } else {
                hi = m - 1
            }
        }
        return hi.toInt()
    }
}
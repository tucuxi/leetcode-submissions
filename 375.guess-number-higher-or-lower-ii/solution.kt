class Solution {
    fun getMoneyAmount(n: Int): Int {
        val dp = Array(n + 1) { IntArray(n + 1) }
        for (len in 2..n) {
            for (start in 1 .. n - len + 1) {
                val end = start + len - 1
                dp[start][end] = (start until end).minOf { i -> i + maxOf(dp[start][i-1], dp[i+1][end]) }
            }
        }
        return dp[1][n]
    }
}
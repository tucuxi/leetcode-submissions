class Solution {
    fun numSquares(n: Int): Int {
        val perfectSquares = makePerfectSquares(n)
        val dp = IntArray(n + 1)
        for (i in 1..n) {
            dp[i] = n + 1
            perfectSquares.forEach {
                if (it <= i) {
                    dp[i] = minOf(dp[i], dp[i - it] + 1)
                }
            }
        }
        return dp[n]
    }

    fun makePerfectSquares(n: Int): List<Int> {
        val res = mutableListOf<Int>()
        var k = 1
        while (k * k <= n) {
            res += k * k
            k++
        }
        return res
    }
}
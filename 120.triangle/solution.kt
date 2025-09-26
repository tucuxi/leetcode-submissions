class Solution {
    fun minimumTotal(triangle: List<List<Int>>): Int {
        val n = triangle.size
        val dp = IntArray(n) { i -> triangle[n - 1][i] }
        for (i in n - 2 downTo 0) {
            for (j in 0..i) {
                dp[j] = triangle[i][j] + minOf(dp[j], dp[j + 1])
            }
        }
        return dp[0]
    }
}
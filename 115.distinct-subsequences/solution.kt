class Solution {
    fun numDistinct(s: String, t: String): Int {
        val dp = Array(t.length + 1) { IntArray(s.length + 1) }
        for (j in 0..s.length) { dp[0][j] = 1 }
        for (i in 0 until t.length) {
            for (j in 0 until s.length) {
                dp[i + 1][j + 1] = if (t[i] == s[j]) {
                    dp[i][j] + dp[i + 1][j]
                } else {
                    dp[i + 1][j]
                }
            }
        }
        return dp[t.length][s.length]
    }
}

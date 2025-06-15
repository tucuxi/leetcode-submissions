class Solution {
    fun numDecodings(s: String): Int {
        val n = s.length
        val dp = IntArray(n + 1)
        dp[n] = 1
        for (i in n - 1 downTo 0) {
            val ch = s[i]
            if (ch != '0') {
                dp[i] = dp[i + 1]
                if (i < n - 1 && (ch == '1' || ch == '2' && s[i + 1] <= '6')) {
                    dp[i] += dp[i + 2]
                }
            }
        }
        return dp[0]
    }
}
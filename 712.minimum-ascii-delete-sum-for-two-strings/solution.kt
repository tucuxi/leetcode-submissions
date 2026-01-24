class Solution {
    fun minimumDeleteSum(s1: String, s2: String): Int {
        val dp = Array(s1.length + 1) { IntArray(s2.length + 1) }
        s1.forEachIndexed { i1, ch1 ->
            s2.forEachIndexed { i2, ch2 ->
                dp[i1 + 1][i2 + 1] = if (ch1 == ch2) {
                    dp[i1][i2] + ch1.code
                } else {
                    maxOf(dp[i1][i2 + 1], dp[i1 + 1][i2])
                }
            }
        }
        return s1.sumOf { it.code } + s2.sumOf { it.code } - 2 * dp[s1.length][s2.length]
    }
}
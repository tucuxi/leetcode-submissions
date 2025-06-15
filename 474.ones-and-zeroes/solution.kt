class Solution {
    fun findMaxForm(strs: Array<String>, m: Int, n: Int): Int {
        val dp = Array(m + 1) { IntArray(n + 1) { 0 } }
        val counts = strs.map { s -> s.count { it == '1' } to s.count { it == '0' } }
        
        counts.forEach { (ones, zeros) ->
            for (i in m downTo zeros) {
                for (j in n downTo ones) {
                    dp[i][j] = maxOf(dp[i][j], 1 + dp[i - zeros][j - ones])
                }
            }
        }
        
        return dp[m][n]
    }
}

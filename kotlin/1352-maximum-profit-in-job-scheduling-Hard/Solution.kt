class Solution {
    fun jobScheduling(startTime: IntArray, endTime: IntArray, profit: IntArray): Int {
        val n = startTime.size
        val timeOrder = startTime.indices.sortedBy { endTime[it] }.toIntArray()

        fun startTime(job: Int) = startTime[timeOrder[job]]
        fun endTime(job: Int) = endTime[timeOrder[job]]
        fun profit(job: Int) = profit[timeOrder[job]]

        val dp = IntArray(n)
        dp[0] = profit(0)
        for (i in 1 until n) {
            dp[i] = maxOf(dp[i - 1], profit(i))
            for (j in i-1 downTo 0) {
                if (endTime(j) <= startTime(i)) {
                    dp[i] = maxOf(dp[i], dp[j] + profit(i))
                    break
                }
            }
        }
        return dp[n - 1]
    }
}
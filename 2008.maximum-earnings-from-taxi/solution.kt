class Solution {
    fun maxTaxiEarnings(n: Int, rides: Array<IntArray>): Long {
        val dp = TreeMap<Int, Long>()
        var maxSum = 0L

        rides.sortBy { it[1] }
        rides.forEach { (start, end, tip) ->
            val priorRideSum = dp.floorEntry(start)?.value ?: 0L
            val afterRideSum = priorRideSum + end - start + tip
            maxSum = maxOf(maxSum, afterRideSum)
            dp[end] = maxSum
        }

        return maxSum
    }
}
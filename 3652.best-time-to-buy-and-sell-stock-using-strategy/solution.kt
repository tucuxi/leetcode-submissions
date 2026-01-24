class Solution {
    fun maxProfit(prices: IntArray, strategy: IntArray, k: Int): Long {
        val halfK = k / 2
        var windowSum = 0L
        var maxWindowSum = 0L
        var total = 0L
        prices.indices.forEach { i ->
            total += prices[i] * strategy[i]
            windowSum += prices[i] * (1 - strategy[i])
            if (i < k - 1) {
                if (i >= halfK - 1) {
                    windowSum -= prices[i - halfK + 1]
                }
            } else {
                maxWindowSum = maxOf(maxWindowSum, windowSum)
                windowSum -= prices[i - halfK + 1] - prices[i - k + 1] * strategy[i - k + 1]
            }
        }
        return total + maxWindowSum
    }
}
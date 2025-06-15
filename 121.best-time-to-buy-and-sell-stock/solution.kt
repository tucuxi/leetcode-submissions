class Solution {
    fun maxProfit(prices: IntArray): Int {
        if (prices.isEmpty()) return 0
        var maxProfitYet = 0
        var lastBuyPrice = prices[0]
        for (i in 1 until prices.size) {
            val profit = prices[i] - lastBuyPrice
            if (profit > maxProfitYet)
                maxProfitYet = profit
            if (prices[i] < lastBuyPrice)
                lastBuyPrice = prices[i]
        }
        return maxProfitYet
    }
}
class Solution {
    fun maxProfit(prices: IntArray): Int {
        fun trade(buy: Int, sell: Int) = prices[sell] - prices[buy]
        fun priceDrop(day: Int) = day > 0 && prices[day] < prices[day - 1]

        var totalProfit = 0
        var buy = 0
        for (day in prices.indices) {
            if (priceDrop(day)) {
                totalProfit += trade(buy, day - 1)
                buy = day
            }
        }
        if (prices[buy] < prices[prices.lastIndex]) {
            totalProfit += trade(buy, prices.lastIndex)
        }
        return totalProfit
    }
}
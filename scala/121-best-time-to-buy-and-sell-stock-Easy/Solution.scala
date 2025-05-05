object Solution {
    def maxProfit(prices: Array[Int]): Int = {
        val futurePrices = new Array[Int](prices.length)
        var maxFuturePrice = 0
        for (i <- prices.length - 1 to 0 by -1) {
            if (prices(i) > maxFuturePrice) maxFuturePrice = prices(i)
            futurePrices(i) = maxFuturePrice
        }
        var maxProfit = 0
        for (i <- 0 until prices.length) { 
            val profit = futurePrices(i) - prices(i)
            if (profit > maxProfit) maxProfit = profit
        }
        maxProfit
    }
}
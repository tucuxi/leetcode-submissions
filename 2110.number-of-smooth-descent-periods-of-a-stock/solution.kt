class Solution {
    fun getDescentPeriods(prices: IntArray): Long {
        var previousPrice = -1
        var previousStart = -1
        var res = 0L

        prices.forEachIndexed { i, price ->
            if (price != previousPrice - 1) {
                previousStart = i
            }
            res += i - previousStart + 1
            previousPrice = price
        }

        return res
    }
}
class Solution {
    fun finalPrices(prices: IntArray): IntArray {
        return prices.mapIndexed { index, price -> price - discount(prices, index) }.toIntArray()
    }

    private fun discount(prices: IntArray, index: Int): Int {
        for (j in index + 1 .. prices.lastIndex) {
            if (prices[j] <= prices[index]) {
                return prices[j]
            }
        }
        return 0
    }
}
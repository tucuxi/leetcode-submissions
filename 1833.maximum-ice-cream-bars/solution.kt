class Solution {
    fun maxIceCream(costs: IntArray, coins: Int): Int {
        val maxCost = costs.max()
        val counts = IntArray(maxCost + 1)

        for (cost in costs) {
            counts[cost]++
        }

        var res = 0
        var rest = coins

        for (cost in 1..maxCost) {
            if (cost > rest) {
                break
            }
            if (counts[cost] == 0) {
                continue
            }
            val bars = minOf(counts[cost], rest / cost)
            res += bars
            rest -= bars * cost
        }

        return res
    }
}
class Solution {
    fun minCost(basket1: IntArray, basket2: IntArray): Long {
        val m = minOf(basket1.min(), basket2.min())

        val count = mutableMapOf<Int, Int>()
        basket1.forEach { f -> count[f] = count.getOrDefault(f, 0) + 1 }
        basket2.forEach { f -> count[f] = count.getOrDefault(f, 0) - 1 }

        val merge = count
            .flatMap { (f, c) ->
                if (c % 2 != 0) return -1
                List<Int>(abs(c) / 2) { f }
            }
            .sorted()

        return merge
            .take(merge.size / 2)
            .fold(0L) { acc, f -> acc + minOf(2 * m, f) }
    }
}
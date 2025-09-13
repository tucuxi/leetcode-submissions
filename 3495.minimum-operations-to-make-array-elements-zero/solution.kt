class Solution {
    fun minOperations(queries: Array<IntArray>): Long {
        return queries.sumOf { (l, r) -> (f(r) - f(l - 1) + 1) / 2 }
    }

    fun f(n: Int): Long {
        var count = 0L
        var i = 1L
        var base = 1

        while (base <= n) {
            val end = minOf(n, 2 * base - 1)
            count += (i + 1) / 2 * (end - base + 1)
            i++
            base *= 2
        }
        return count
    }
}
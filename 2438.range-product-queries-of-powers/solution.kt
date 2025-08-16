const val MOD = 1_000_000_007

class Solution {
    fun productQueries(n: Int, queries: Array<IntArray>): IntArray {
        val powers = (0..30).map { 1 shl it % MOD }.filter { it and n != 0 }
        val res = queries.map { (l, r) -> 
            powers.subList(l, r + 1).fold(1L) { acc, p -> acc * p % MOD }.toInt()
        }
        return res.toIntArray()
    }
}
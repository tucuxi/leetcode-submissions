class Solution {
    fun getKth(lo: Int, hi: Int, k: Int): Int {
        val memo = mutableMapOf(1 to 0)

        fun power(x: Int): Int =
            memo.getOrElse(x) {          
                val p = 1 + if (x % 2 == 0) {
                    power(x / 2)
                } else {
                    power(3 * x + 1)
                }
                memo[x] = p
                return p
            }

       return (lo..hi)
            .map { x -> x to power(x) }
            .sortedBy { (_, p) -> p }
            .get(k - 1)
            .first
    }
}
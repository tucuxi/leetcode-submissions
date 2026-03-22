class Solution {
    fun countCommas(n: Long): Long {
        var res = 0L
        var limit = 1_000_000_000_000_000L
        var n = n

        for (commas in 5 downTo 1) {
            if (n >= limit) {
                res += commas * (n - limit + 1)
                n = limit - 1
            }
            limit /= 1000
        }

        return res
    }
}
class Solution {
    fun minFlips(s: String): Int {
        val n = s.length
        var k = 0
        var p = 0

        for (c in s) {
            k += c.toInt() and 1 xor p
            p = 1 - p
        }

        var res = minOf(k, n - k)

        if (n % 2 == 0) return res

        var b = 0
        p = 0

        for (c in s) {
            k += (c.toInt() and 1 xor p xor 1) - (c.toInt() and 1 xor p)
            p = 1 - p
            res = minOf(res, k, n - k)
        }
 
        return res
    }
}
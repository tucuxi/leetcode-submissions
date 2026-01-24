class Solution {
    fun findNthDigit(n: Int): Int {
        var l = 1
        var r = Int.MAX_VALUE

        while (l < r) {
            val m = l + (r - l) / 2
            if (totalNumberOfDigits(m) < n) {
                l = m + 1
            } else {
                r = m
            }
        }

        return digit(l, (totalNumberOfDigits(l) - n).toInt())
    }

    fun totalNumberOfDigits(n: Int): Long {
        return limits.sumOf { l -> maxOf(0, n.toLong() - l) }
    }

    fun digit(n: Int, i: Int): Int {
        val s = n.toString()
        return (s[s.length - 1 - i] - '0').toInt()
    }

    companion object {
        val limits = listOf(0L, 9, 99, 999, 9999, 99999, 999999, 9999999, 99999999)
    }
}
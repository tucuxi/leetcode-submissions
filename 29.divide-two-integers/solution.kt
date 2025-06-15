import kotlin.math.abs

class Solution {
    fun divide(dividend: Int, divisor: Int): Int {
        if (divisor == 1) return dividend
        if (dividend == Int.MIN_VALUE && divisor == -1) return Int.MAX_VALUE
        if (dividend == Int.MIN_VALUE && divisor == Int.MIN_VALUE) return 1
        var d = abs(divisor)
        var dd = abs(dividend)
        var r = 0
        while (dd - d >= 0) {
            var i = 0
            while (dd - (d shl i) >= 0 && d shl i > 0) i++
            i--
            r += 1 shl i
            dd -= d shl i
        }
        return if (dividend xor divisor >= 0) r else -r
    }
}
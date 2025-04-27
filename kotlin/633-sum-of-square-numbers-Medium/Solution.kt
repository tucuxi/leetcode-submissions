import kotlin.math.sign
import kotlin.math.sqrt

class Solution {
    fun judgeSquareSum(c: Int): Boolean {
        var a = 0
        var b = sqrt(c.toDouble()).toInt()
        while (a <= b) {
            when ((c - sq(a) - sq(b)).sign) {
                -1 -> b--
                0 -> return true
                1 -> a++
            }
        }
        return false
    }
}

inline fun sq(a: Int) = a * a
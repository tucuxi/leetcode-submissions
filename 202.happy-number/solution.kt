class Solution {
    fun isHappy(n: Int): Boolean {
        val s = mutableSetOf<Int>()
        var p = n
        while (p != 1) {
            var q = 0
            while (p > 0) {
                q += square(p % 10)
                p /= 10
            }
            if (!s.add(q)) {
                return false
            }
            p = q
        }
        return true
    }
}

inline fun square(a: Int): Int = a * a
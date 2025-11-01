class Solution {
    fun scoreBalance(s: String): Boolean {
        var left = 0
        var right = s.sumOf { it.value() }

        for (c in s) {
            val v = c.value()
            left += v
            right -= v
            if (left >= right) {
                break
            }
        }

        return left == right
    }
}

inline fun Char.value(): Int = this - 'a' + 1
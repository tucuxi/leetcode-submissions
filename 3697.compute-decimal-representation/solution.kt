class Solution {
    fun decimalRepresentation(n: Int): IntArray {
        var rest = n
        var base = 1
        val digits = mutableListOf<Int>()
        while (rest > 0) {
            val d = rest % 10
            if (d > 0) {
                digits.add(d * base)
            }
            rest /= 10
            base *= 10
        }
        return IntArray(digits.size) { digits[digits.lastIndex - it] }
    }
}
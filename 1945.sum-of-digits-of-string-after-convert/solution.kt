class Solution {
    fun getLucky(s: String, k: Int): Int {
        var digits = s.flatMap { charToDigits(it) }
        repeat(k) {
            digits = transform(digits)
        }
        return digitsToInt(digits)
    }
    
    fun charToDigits(ch: Char): List<Int> {
        val num = ch - 'a' + 1
        return if (num < 10) listOf(num) else listOf(num / 10, num % 10)
    }
    
    fun transform(s: List<Int>): List<Int> {
        var sum = s.reduce { a, b -> a + b }
        val res = mutableListOf<Int>()
        while (sum > 0) {
            res.add(0, sum % 10)
            sum /= 10
        }
        return res
    }
    
    fun digitsToInt(digits: List<Int>) =
        digits.fold(0) { acc, digit -> acc * 10 + digit }
}

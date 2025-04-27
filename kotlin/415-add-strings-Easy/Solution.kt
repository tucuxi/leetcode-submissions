class Solution {
    fun addStrings(num1: String, num2: String): String {
        val zero = '0'.toInt()
        with (StringBuilder()) {
            var carry = 0
            for (i in 0 until maxOf(num1.length, num2.length)) {
                val a = num1.getDigit(num1.lastIndex - i) + num2.getDigit(num2.lastIndex - i) + carry
                append((a % 10 + zero).toChar())
                carry = a / 10
            }
            if (carry > 0) {
                append((carry + zero).toChar())
            }
            return reverse().toString()
        }
    }
    
    fun String.getDigit(index: Int): Int = (getOrElse(index) { '0' } - '0').toInt()
}
class Solution {
    fun plusOne(digits: IntArray) =
        sequence {
            val carry = digits.foldRight(1) { carry, digit ->
                val sum = digit + carry
                yield(sum % 10)
                sum / 10
            }
            if (carry > 0)
                yield(carry)
        }.toList().reversed().toIntArray()
}
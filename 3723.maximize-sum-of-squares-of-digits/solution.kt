class Solution {
    fun maxSumOfSquares(num: Int, sum: Int): String {
        var rest = sum
        val result = ByteArray(num)

        repeat(num) { i ->
            val digit = minOf(rest, 9)
            rest -= digit
            result[i] = (digit + '0'.code).toByte()
        }

        return if (rest > 0) {
            ""
        } else {
            String(result)
        }
    }
}
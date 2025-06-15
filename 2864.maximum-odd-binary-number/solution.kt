class Solution {
    fun maximumOddBinaryNumber(s: String): String {
        val h = s.groupingBy { it }.eachCount()
        val ones = h.getOrElse('1') { throw IllegalArgumentException("s must contain at least one '1'") }
        val zeros = h.getOrDefault('0', 0)
        return "1".repeat(ones - 1) + "0".repeat(zeros) + "1"
    }
}
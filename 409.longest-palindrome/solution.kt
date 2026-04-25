class Solution {
    fun longestPalindrome(s: String): Int {
        return minOf(
            s.length,
            1 + s
                .groupingBy { it }
                .eachCount()
                .values
                .sumOf { it / 2 * 2 }
        )
    }
}
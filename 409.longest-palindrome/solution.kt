class Solution {
    fun longestPalindrome(s: String): Int {
        return minOf(
            s.length,
            1 + s
                .groupingBy { it }
                .eachCount()
                .map { (_, v) -> v / 2 * 2 }
                .sum()
        )
    }
}
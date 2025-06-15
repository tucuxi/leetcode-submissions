class Solution {
    fun generateTheString(n: Int): String {
        val p = n + n % 2 - 1
        return "p".repeat(p) + "q".repeat(n - p)
    }
}
class Solution {
    fun minCut(s: String): Int {
        val n = s.length
        val cuts = IntArray(n) { it }
        val palindrome = Array(n) { BooleanArray(n) }
        for (i in 0 until n) {
            for (j in 0..i) {
                if (s[i] == s[j] && (i - j < 2 || palindrome[j + 1][i - 1])) {
                    palindrome[j][i] = true
                    cuts[i] = if (j == 0) 0 else minOf(cuts[i], cuts[j - 1] + 1)
                }
            }
        }
        return cuts[n - 1]
    }
}
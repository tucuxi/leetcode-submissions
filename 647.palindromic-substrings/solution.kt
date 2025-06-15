class Solution {
    fun countSubstrings(s: String): Int {
        val N = s.length
        val p = Array(N) { BooleanArray(N) }
        for (i in N-1 downTo 0) {
            for (j in i until N) {
                p[i][j] = s[i] == s[j] && (i >= j - 1 || p[i + 1][j - 1])
            }
        }
        return p.map { q -> q.count { it } }.sum()
    }
}
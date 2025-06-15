class Solution {
    fun findValidPair(s: String): String {
        val f = s.groupingBy { it }.eachCount()
        for (i in 0 until s.lastIndex) {
            if (s[i] != s[i + 1] && f[s[i]] == s[i] - '0' && f[s[i + 1]] == s[i + 1] - '0') {
                return s.substring(i .. i + 1)
            }
        }
        return ""
    }
}
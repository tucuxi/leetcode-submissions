class Solution {
    fun isSubsequence(s: String, t: String): Boolean {
        var p = 0
        var q = 0
        while (p < s.length && q < t.length) {
            if (s[p] == t[q++]) {
                p++
            }
        }
        return p == s.length
    }
}
class Solution {
    fun countBinarySubstrings(s: String): Int {
        var res = 0
        var p = 1
        var q = 0

        for (i in 1 until s.length) {
            if (s[i] == s[i - 1]) {
                p++
            } else {
                q = p
                p = 1
            }
            if (q >= p) {
                res++
            }
        }
        return res
    }
}
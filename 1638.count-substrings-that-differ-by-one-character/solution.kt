class Solution {
    fun countSubstrings(s: String, t: String): Int {
        var res = 0
        for (i in s.indices) {
            for (j in t.indices) {
                var diffs = 0
                var k = 0
                val l = minOf(s.length - i, t.length - j)
                while (k < l && diffs <= 1) {
                    if (s[i + k] != t[j + k]) {
                        diffs++
                    }
                    if (diffs == 1) {
                        res++
                    }
                    k++
                }
            }
        }
        return res
    }
}
class Solution {
    fun countValidPrefixes(s: String): Int {
        var diff = 0

        return s.count { ch ->
            if (ch == '0') {
                diff--
            } else {
                diff++
            }
            -1 <= diff && diff <= 1
        }
    }
}
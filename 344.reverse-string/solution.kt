class Solution {
    fun reverseString(s: CharArray): Unit {
        var p = 0
        var q = s.lastIndex
        while (p < q) {
            val h = s[p]
            s[p++] = s[q]
            s[q--] = h
        }
    }
}
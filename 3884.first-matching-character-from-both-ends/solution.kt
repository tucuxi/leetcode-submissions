class Solution {
    fun firstMatchingIndex(s: String): Int {
        var i = 0
        var j = s.lastIndex
        while (i <= j) {
            if (s[i] == s[j]) return i
            i++
            j--
        }
        return -1
    }
}
class Solution {
    fun minimumMoves(s: String): Int {
        var res = 0
        var i = 0
        while (i < s.length) {
            if (s[i] == 'X') {
                res++
                i += 3
            } else {
                i++
            }
        }
        return res
    }
}
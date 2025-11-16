class Solution {
    fun maxOperations(s: String): Int {
        var res = 0
        var ones = 0
        var i = 0

        while (i < s.length) {
            if (s[i] == '0') {
                while (i + 1 < s.length && s[i + 1] == '0') {
                    i++
                }
                res += ones
            } else {
                ones++
            }
            i++
        }

        return res
    }
}
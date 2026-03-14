class Solution {
    fun maximumXor(s: String, t: String): String {
        var res = StringBuilder(s)
        var t1 = t.count { it == '1' }
        var i = 0

        while (t1 > 0 && i < s.length) {
            if (s[i] == '0') {
                res[i] = '1'
                t1--
            }
            i++
        }

        while (t1 > 0) {
            if (s[--i] == '1') {
                res[i] = '0'
                t1--
            }
        }

        return res.toString()        
    }
}
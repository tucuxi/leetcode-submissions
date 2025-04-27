class Solution {
    fun findLUSlength(strs: Array<String>): Int {
        var res = -1
        loop@ for (i in 0 until strs.size) {
            for (j in 0 until strs.size) {
                if (i == j) continue
                if (isSubsequence(strs[i], strs[j])) continue@loop        
            }
            res = maxOf(res, strs[i].length)
        }
        return res
    }
    
    fun isSubsequence(s: String, t: String): Boolean {
        var i = 0
        var j = 0
        while (i < s.length && j < t.length) {
            if (s[i] == t[j]) i++
            j++
        }
        return i == s.length
    }
}
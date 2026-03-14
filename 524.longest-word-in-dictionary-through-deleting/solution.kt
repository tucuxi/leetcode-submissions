class Solution {
    fun findLongestWord(s: String, dictionary: List<String>): String {
        var res = ""
        
        for (d in dictionary) {
            var i = 0
            var j = 0
            
            while (i < s.length && j < d.length) {
                if (s[i] == d[j]) {
                    j++
                }
                i++
            }
            if (j == d.length) {
                when {
                    d.length > res.length -> res = d
                    d.length == res.length -> res = minOf(res, d)
                }
            }
        }
        
        return res
    }
}
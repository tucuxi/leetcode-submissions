class Solution {
    fun minimumDeletions(s: String): Int {
        var res = s.length
        var a = s.count { ch -> ch == 'a' }
        var b = 0
        s.forEach { ch ->
            if (ch == 'a') {
                a--
            }
            res = minOf(res, a + b)
            if (ch == 'b') {
                b++
            }
        }
        return res
    }
}
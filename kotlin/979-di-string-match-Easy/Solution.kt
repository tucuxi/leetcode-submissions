class Solution {
    fun diStringMatch(s: String): IntArray {
        var p = 0
        var q = s.length
        val res = mutableListOf<Int>()
        s.forEach { c ->
            if (c == 'I') {
                res.add(p++)
            } else {
                res.add(q--)
            }
        }
        res.add(p)
        return res.toIntArray()
    }
}
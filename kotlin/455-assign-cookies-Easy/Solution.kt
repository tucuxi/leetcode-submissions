class Solution {
    fun findContentChildren(g: IntArray, s: IntArray): Int {
        val gs = g.sorted()
        val ss = s.sorted()
        var res = 0
        var i = 0
        var j = 0
        while (i < gs.size && j < ss.size) {
            if (gs[i] <= ss[j]) {
                res++
                i++
            }
            j++
        }
        return res
    }
}
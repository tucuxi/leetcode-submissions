class Solution {
    fun binaryGap(n: Int): Int {
        var res = 0
        val s = n.toString(2)
        var p = s.indexOf('1')
        while (p != -1) {
            val q = s.indexOf('1', p + 1)
            res = maxOf(res, q - p)
            p = q
        }
        return res
    }
}
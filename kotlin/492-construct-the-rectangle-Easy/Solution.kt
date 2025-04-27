class Solution {
    fun constructRectangle(area: Int): IntArray {
        val res = intArrayOf(0, 0)
        var w = 1
        while (w * w <= area) {
            if (area % w == 0) {
                res[0] = area / w
                res[1] = w
            }
            w++
        }
        return res
    }
}
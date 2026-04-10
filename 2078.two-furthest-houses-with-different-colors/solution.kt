class Solution {
    fun maxDistance(colors: IntArray): Int {
        var res = 0

        for (i in colors.indices) {
            if (colors[i] != colors[0]) {
                res = maxOf(res, i)
            }
            if (colors[i] != colors[colors.lastIndex]) {
                res = maxOf(res, colors.lastIndex - i)
            }
        }

        return res
    }
}
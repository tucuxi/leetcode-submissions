class Solution {
    fun maxDistance(colors: IntArray): Int {
        var res = 0
        val cl = colors[0]
        val cr = colors[colors.lastIndex]

        colors.asIterable().forEachIndexed { i, c ->
            if (c != cl) {
                res = maxOf(res, i)
            }
            if (c != cr) {
                res = maxOf(res, colors.lastIndex - i)
            }
        }

        return res
    }
}
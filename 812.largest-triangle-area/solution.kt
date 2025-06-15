import kotlin.math.abs

class Solution {
    fun largestTriangleArea(points: Array<IntArray>): Double {
        var maxArea = 0.0
        for (i in 0 .. points.lastIndex) {
            val x1 = points[i][0].toDouble()
            val y1 = points[i][1].toDouble()
            for (j in i + 1 .. points.lastIndex) {
                val x2 = points[j][0].toDouble()
                val y2 = points[j][1].toDouble()
                for (k in j + 1 .. points.lastIndex) {
                    val x3 = points[k][0].toDouble()
                    val y3 = points[k][1].toDouble()
                    maxArea = maxOf(maxArea, abs(0.5 * (x2 * y3 + x1 * y2 + x3 * y1 - x3 * y2 - x2 * y1 - x1 * y3)))
                }
            }
        }
        return maxArea
    }
}
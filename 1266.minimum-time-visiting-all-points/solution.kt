class Solution {
    fun minTimeToVisitAllPoints(points: Array<IntArray>): Int {
        return (1 until points.size).sumOf { i ->
            val dx = abs(points[i - 1][0] - points[i][0])
            val dy = abs(points[i - 1][1] - points[i][1])
            maxOf(dx, dy)
        }
    }
}
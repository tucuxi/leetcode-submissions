class Solution {
    fun minTimeToVisitAllPoints(points: Array<IntArray>): Int {
        var res = 0
        for (i in 1 until points.size) {
            var dx = Math.abs(points[i][0] - points[i - 1][0])
            var dy = Math.abs(points[i][1] - points[i - 1][1])
            res += Math.max(dx, dy)
        }
        return res
    }
}
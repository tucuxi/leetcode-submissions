import kotlin.math.abs

class Solution {
    fun nearestValidPoint(x: Int, y: Int, points: Array<IntArray>): Int {
        var minDistance = Int.MAX_VALUE
        var minIndex = -1
        points.forEachIndexed { index, (xp, yp) ->
            if (x == xp || y == yp) {
                val dist = abs(x - xp) + abs(y -yp)
                if (dist < minDistance) {
                    minDistance = dist
                    minIndex = index
                }
            }
        }
        return minIndex
    }
}
class Solution {
    fun numberOfPairs(points: Array<IntArray>): Int {
        points.sortWith(
            compareBy<IntArray> { it[0] }
                .thenByDescending<IntArray> { it[1] }
        )
        var res = 0
        points.forEachIndexed { i, (_, y1) ->
            var l = Int.MIN_VALUE
            var h = y1
            points.drop(i + 1).forEach { (x2, y2) ->
                if (y2 <= h && y2 > l) {
                    res++
                    l = y2
                    if (y2 == h) {
                        h--
                    }
                }
            }
        }
        return res
    }
}
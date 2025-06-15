class Solution {
    fun countPoints(points: Array<IntArray>, queries: Array<IntArray>): IntArray =
        queries.map { (x, y, r) ->
            points.count { (xp, yp) ->
                square(x - xp) + square(y - yp) <= square(r)
            }
        }.toIntArray()
}

fun square(a: Int) = a * a
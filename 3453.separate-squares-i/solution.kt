const val EPSILON = 1e-5
class Solution {
    fun separateSquares(squares: Array<IntArray>): Double {
        var l = 0.0
        var r = 1e9
        while (r - l > EPSILON) {
            val m = (l + r) / 2
            var lower = 0.0
            var upper = 0.0
            squares.forEach { (_, y, l) ->
                val area = l.toDouble() * l.toDouble()
                when {
                    m <= y -> {
                        upper += area
                    }
                    m >= y + l -> {
                        lower += area
                    }
                    else -> {
                        val r = (m - y) / l
                        upper += area * (1 -r)
                        lower += area * r
                    }
                }
            }
            if (lower < upper) {
                l = m
            } else {
                r = m
            }
        }
        return l
    }
}
object Solution {
    def checkStraightLine(coordinates: Array[Array[Int]]): Boolean = {
        val x0 = coordinates(0)(0)
        val y0 = coordinates(0)(1)
        val dx1 = coordinates(1)(0) - x0
        val dy1 = coordinates(1)(1) - y0
        for (p <- coordinates.tail) {
            val dxi = p(0) - x0
            val dyi = p(1) - y0
            if (dx1 * dyi != dy1 * dxi) return false
        }
        true
    }
}
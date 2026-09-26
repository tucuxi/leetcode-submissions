class Solution {
    fun checkOverlap(radius: Int, xCenter: Int, yCenter: Int, x1: Int, y1: Int, x2: Int, y2: Int): Boolean {
        val dx = when {
            xCenter < x1 -> xCenter - x1
            xCenter > x2 -> xCenter - x2
            else -> 0
        }
        val dy = when {
            yCenter < y1 -> yCenter - y1
            yCenter > y2 -> yCenter - y2
            else -> 0
        }
        return dx * dx + dy * dy <= radius * radius
    }
}
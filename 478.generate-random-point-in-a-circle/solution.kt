class Solution(val radius: Double, val x_center: Double, val y_center: Double) {

    fun randPoint(): DoubleArray {
        val r = radius * sqrt(Random.nextDouble(1.0))
        val phi = Random.nextDouble(2 * PI)
        return doubleArrayOf(x_center + r * cos(phi), y_center + r * sin(phi))
    }

}

/**
 * Your Solution object will be instantiated and called as such:
 * var obj = Solution(radius, x_center, y_center)
 * var param_1 = obj.randPoint()
 */
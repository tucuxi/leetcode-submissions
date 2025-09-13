class Solution {
    fun findClosest(x: Int, y: Int, z: Int): Int {
        val dx = abs(x - z)
        val dy = abs(y - z)
        return when {
            dx < dy -> 1
            dx > dy -> 2
            else -> 0
        }
    }
}
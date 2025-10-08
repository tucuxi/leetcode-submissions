class Solution {
    fun swimInWater(grid: Array<IntArray>): Int {
        val n = grid.size
        val distance = Array(n) { IntArray(n) { Int.MAX_VALUE } }
        distance[0][0] = grid[0][0]
        val queue = mutableListOf(0 to 0)
        while (queue.isNotEmpty()) {
            val (r, c) = queue.removeFirst()
            STEPS
                .map { (dr, dc) -> r + dr to c + dc }
                .filter { (r2, c2) -> r2 in 0 until n && c2 in 0 until n }
                .forEach { (r2, c2) ->
                    val d = maxOf(distance[r][c], grid[r2][c2])
                    if (d < distance[r2][c2]) {
                        distance[r2][c2] = d
                        queue.add(r2 to c2)
                    }
                }
        }
        return distance[n - 1][n - 1]
    }

    companion object {
        private val STEPS = listOf(-1 to 0, 0 to 1, 1 to 0, 0 to -1)
    }
}
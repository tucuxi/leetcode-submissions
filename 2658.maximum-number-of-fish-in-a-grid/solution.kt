class Solution {
    fun findMaxFish(grid: Array<IntArray>): Int {
        val m = grid.indices
        val n = grid.first().indices

        fun dfs(r: Int, c: Int): Int {
            return if (r !in m || c !in n || grid[r][c] == 0) {
                0
            } else {
                val f = grid[r][c].also { grid[r][c] = 0 }
                f + listOf(r - 1 to c, r + 1 to c, r to c - 1, r to c + 1)
                    .fold(grid[r][c]) { acc, (r2, c2) -> acc + dfs(r2, c2) }
            }
        }

        return m.flatMap { r -> n.map { c -> dfs(r, c) } }.max()
    }
}
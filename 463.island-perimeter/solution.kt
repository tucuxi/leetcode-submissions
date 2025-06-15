import kotlin.math.abs

class Solution {
    fun islandPerimeter(grid: Array<IntArray>): Int {
        val lastRow = grid.size - 1
        val lastCol = grid[0].size - 1
        var p = 0
        
        fun leftOf(i: Int, j: Int) = if (j == 0) 0 else grid[i][j - 1]
        fun above(i: Int, j: Int) = if (i == 0) 0 else grid[i - 1][j]
        
        for (i in 0..lastRow) {
            for (j in 0..lastCol) {
                p += abs(grid[i][j] - above(i, j)) + abs(grid[i][j] - leftOf(i, j))
            }
        }
        for (i in 0..lastRow) {
            p += grid[i][lastCol]
        }
        for (j in 0..lastCol) {
            p += grid[lastRow][j]
        }
        return p
    }
}
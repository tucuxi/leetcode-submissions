class Solution {
    fun satisfiesConditions(grid: Array<IntArray>): Boolean {
        return grid.indices.all { i ->
            grid[i].withIndex().all { (j, value) ->
                (i == grid.lastIndex || value == grid[i+1][j]) &&
                (j == grid[i].lastIndex || value != grid[i][j+1])
            }
        }
    }
}
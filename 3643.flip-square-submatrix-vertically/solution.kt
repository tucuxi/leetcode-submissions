class Solution {
    fun reverseSubmatrix(grid: Array<IntArray>, x: Int, y: Int, k: Int): Array<IntArray> {
        var r1 = x
        var r2 = x + k - 1
        while (r1 < r2) {       
            for (c in y until y + k) {
                grid[r1][c] = grid[r2][c].also { grid[r2][c] = grid[r1][c] }
            }
            r1++
            r2--
        }
        return grid
    }
}
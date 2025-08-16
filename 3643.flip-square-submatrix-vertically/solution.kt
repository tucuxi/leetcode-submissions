class Solution {
    fun reverseSubmatrix(grid: Array<IntArray>, x: Int, y: Int, k: Int): Array<IntArray> {
        for (r in x until x + k / 2) {
            for (c in y until y + k) {
                val r2 = x + k - 1 - (r - x)
                val h = grid[r][c]
                grid[r][c] = grid[r2][c]
                grid[r2][c] = h
            }
        }
        return grid
    }
}
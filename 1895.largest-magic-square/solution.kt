class Solution {
    fun largestMagicSquare(grid: Array<IntArray>): Int {
        val m = grid.size
        val n = grid[0].size
        val rows = Array(m + 1) { IntArray(n + 1) }
        val cols = Array(m + 1) { IntArray(n + 1) }
        val d1 = Array(m + 1) { IntArray(n + 1) }
        val d2 = Array(m + 1) { IntArray(n + 2) } // Extra padding for anti-diag

        for (i in 0 until m) {
            for (j in 0 until n) {
                rows[i + 1][j + 1] = rows[i + 1][j] + grid[i][j]
                cols[i + 1][j + 1] = cols[i][j + 1] + grid[i][j]
                d1[i + 1][j + 1] = d1[i][j] + grid[i][j]
                d2[i + 1][j + 1] = d2[i][j + 2] + grid[i][j]
            }
        }

        for (side in minOf(m, n) downTo 2) {
            for (i in 0 .. m - side) {
                columns@ for (j in 0 .. n - side) {
                    val target = rows[i + 1][j + side] - rows[i + 1][j]
                    
                    for (k in 1 until side) {
                        val rSum = rows[i + 1 + k][j + side] - rows[i + 1 + k][j]
                        if (rSum != target) continue@columns
                    }

                    for (k in 0 until side) {
                        val cSum = cols[i + side][j + 1 + k] - cols[i][j + 1 + k]
                        if (cSum != target) continue@columns
                    }

                    val diag1Sum = d1[i + side][j + side] - d1[i][j]
                    if (diag1Sum != target) continue@columns

                    val diag2Sum = d2[i + side][j + 1] - d2[i][j + side + 1]
                    if (diag2Sum != target) continue@columns
                    
                    return side
                }
            }
        }
        return 1
    }
}
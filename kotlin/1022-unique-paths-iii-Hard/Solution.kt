class Solution {
    private val DIR = listOf(Pair(-1, 0), Pair(0, -1), Pair(1, 0), Pair(0, 1))
    
    fun uniquePathsIII(grid: Array<IntArray>): Int {
        val m = grid.size
        val n = grid[0].size
        var res = 0
        
        fun countEmpty() =
            grid.fold(0) { sum, row -> sum + row.count { sq -> sq == 0 }}
        
        fun dfs(row: Int, col: Int, remainingSquares: Int) {
            if (row !in 0 until m || col !in 0 until n) {
                return
            }
            val sq = grid[row][col]
            if (sq == 2) {
                if (remainingSquares == 0) {
                    res++
                }
            } else if (sq == 0) {
                grid[row][col] = 3
                for ((drow, dcol) in DIR) { 
                    dfs(row + drow, col + dcol, remainingSquares - 1)
                }
                grid[row][col] = sq
            }
        }
        
        fun findStart(): Pair<Int, Int> {
            for (row in 0 until m) {
                for (col in 0 until n) {
                    if (grid[row][col] == 1) {
                        return Pair(row, col)
                    }
                }
            }
            return Pair(-1, -1)
        }
        
        val (startRow, startCol) = findStart()
        for ((drow, dcol) in DIR) {
            dfs(startRow + drow, startCol + dcol, countEmpty())
        }
        return res
    }
}
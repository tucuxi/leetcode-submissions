class Solution {
    fun solve(board: Array<CharArray>): Unit {
        val m = board.size
        val n = board[0].size
        
        fun flood(row: Int, col: Int) {
            if (board[row][col] == 'O') {
                board[row][col] = 'R'
                if (row - 1 >= 0) {
                    flood(row - 1, col)
                }
                if (row + 1 < m) {
                    flood(row + 1, col)
                }
                if (col - 1 >= 0) {
                    flood(row, col - 1)
                }
                if (col + 1 < n) {
                    flood(row, col + 1)
                }
            }
        }
        
        for (row in 0 until m) {
            flood(row, 0)
            flood(row, n - 1)
        }
        for (col in 0 until n) {
            flood(0, col)
            flood(m - 1, col)
        }
        for (row in 0 until m) {
            for (col in 0 until n) {
                board[row][col] = if (board[row][col] == 'R') 'O' else 'X'
            }
        }
    }
}
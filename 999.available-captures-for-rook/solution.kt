class Solution {
    fun numRookCaptures(board: Array<CharArray>): Int {
        var res = 0

        fun findRook(): Pair<Int, Int> {
            for (row in 0 .. board.lastIndex) {
                for (col in 0 .. board[row].lastIndex) {
                    if (board[row][col] == 'R') {
                        return Pair(row, col)
                    }
                }
            }
            throw IllegalArgumentException()
        }
        
        val (row, col) = findRook()
        
        var r1 = row
        while (--r1 >= 0) {        
            if (board[r1][col] == 'p') { res++ }
            if (board[r1][col] != '.') { break }
        }
        
        var r2 = row
        while (++r2 <= board.lastIndex) {
            if (board[r2][col] == 'p') { res++ }
            if (board[r2][col] != '.') { break }
        }
        
        var c1 = col
        while (--c1 >= 0) {
            if (board[row][c1] == 'p') { res++ }
            if (board[row][c1] != '.') { break }
        }
        
        var c2 = col
        while (++c2 <= board[col].lastIndex) {
            if (board[row][c2] == 'p') { res++ }
            if (board[row][c2] != '.') { break }
        }
        
        return res
    }
}
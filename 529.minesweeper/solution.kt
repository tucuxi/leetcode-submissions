typealias Board = Array<CharArray>

class Solution {
    fun updateBoard(board: Array<CharArray>, click: IntArray): Array<CharArray> {
        val (row, col) = click
        when (board[row][col]) {
            'M' -> board[row][col] = 'X'
            'E' -> reveal(board, row, col)
        }
        return board
    }
    
    private fun reveal(board: Board, row: Int, col: Int) {
        val a = adjacentMines(board, row, col)       
        if (a > 0) {
            board[row][col] = '0' + a
            return
        }
        board[row][col] = 'B'
        for (r in maxOf(0, row-1) .. minOf(board.lastIndex, row+1)) {
            for (c in maxOf(0, col-1) .. minOf(board[r].lastIndex, col+1)) {
                if (board[r][c] == 'E') {
                    reveal(board, r, c)
                }
            }
        }
    }
    
    private fun adjacentMines(board: Board, row: Int, col: Int): Int {
        var count = 0
        for (r in maxOf(0, row-1) .. minOf(board.lastIndex, row+1)) {
            for (c in maxOf(0, col-1) .. minOf(board[r].lastIndex, col+1)) {
                if (board[r][c] in setOf('M', 'X')) {
                    count++
                }
            }
        }
        return count
    }
}
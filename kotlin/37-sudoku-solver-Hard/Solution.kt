class Solution {
    fun solveSudoku(board: Array<CharArray>) {
        Model(board).solve()
    }
}

class Model(private val board: Array<CharArray>) {
    private val n = board.size
    private val row = Array(n) { BooleanArray(n) }
    private val col = Array(n) { BooleanArray(n) }
    private val box = Array(n) { BooleanArray(n) }
    private val emptyCells = mutableListOf<Pair<Int, Int>>()

    init {
        for (r in 0 until n) {
            for (c in 0 until n) {
                if (board[r][c] == '.') {
                    emptyCells.add(Pair(r, c))
                } else {
                    flip(r, c, board[r][c])
                }
            }
        }
    }

    private fun flip(r: Int, c: Int, ch: Char) {
        val p = ch - '1'
        row[r][p] = !row[r][p]
        col[c][p] = !col[c][p]
        val q = r / 3 * 3 + c / 3
        box[q][p] = !box[q][p]
    }

    private fun valid(r: Int, c: Int, ch: Char): Boolean {
        val p = ch - '1'
        return !row[r][p] && !col[c][p] && !box[r / 3 * 3 + c / 3][p]
    }

    private fun dfs(emptyCellIndex: Int): Boolean {
        if (emptyCellIndex == emptyCells.size) {
            return true
        }
        val (r, c) = emptyCells[emptyCellIndex]
        for (ch in '1'..'9') {
            if (valid(r, c, ch)) {
                flip(r, c, ch)
                if (dfs(emptyCellIndex + 1)) {
                    board[r][c] = ch
                    return true
                }
                flip(r, c, ch)
            }
        }
        return false
    }

    fun solve() = dfs(0)
}
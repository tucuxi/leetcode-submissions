class Solution {
    fun isValidSudoku(board: Array<CharArray>): Boolean {
	    val rows = Array(9) { mutableSetOf<Char>() }
	    val cols = Array(9) { mutableSetOf<Char>() }
	    val boxes = Array(3) { Array(3) { mutableSetOf<Char>() } }

    	for (row in board.indices) {
	    	for (col in board[row].indices) {
		    	with (board[row][col]) {
			        if (this != '.') {
                        if (!rows[row].add(this)
                                || !cols[col].add(this)
                                || !boxes[row / 3][col / 3].add(this)) {
                            return false
                        }
                    }
			    }
	    	}
    	}

	    return true
    }
}
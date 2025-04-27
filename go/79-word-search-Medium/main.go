func exist(board [][]byte, word string) bool {
    var dfs func(row, col, matched int) bool
    dfs = func(row, col, matched int) bool {
        if board[row][col] == word[matched] {
            if matched + 1 == len(word) {
                return true
            }
            board[row][col] = '#'
            if row > 0 && dfs(row - 1, col, matched + 1) {
                return true
            }
            if col > 0 && dfs(row, col - 1, matched + 1) {
                return true
            }
            if row + 1 < len(board) && dfs(row + 1, col, matched + 1) {
                return true
            }
            if col + 1 < len(board[row]) && dfs(row, col + 1, matched + 1) {
                return true
            }
            board[row][col] = word[matched]
        }
        return false
    }

    for row := range board {
        for col := range board[row] {
            if dfs(row, col, 0) {
                return true
            }
        }
    }
    return false
}
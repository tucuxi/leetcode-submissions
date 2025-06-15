func countBattleships(board [][]byte) int {
    res := 0
    for i := range board {
        for j := range board[i] {
            if board[i][j] == 'X' &&
                    (i == 0 || board[i - 1][j] != 'X') &&
                    (j == 0 || board[i][j - 1] != 'X') {
                res++
            }
        }
    }
    return res
}
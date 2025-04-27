func isValidSudoku(board [][]byte) bool {
    rows := [9]map[byte]bool{{}, {}, {}, {}, {}, {}, {}, {}, {}}
    cols := [9]map[byte]bool{{}, {}, {}, {}, {}, {}, {}, {}, {}}
    boxes := [3][3]map[byte]bool{{{}, {}, {}}, {{}, {}, {}}, {{}, {}, {}}}

    for row := range board {
        for col := range board[row] {
            if b := board[row][col]; b != '.' {
                if rows[row][b] || cols[col][b] || boxes[row / 3][col / 3][b] {
                    return false
                }
                rows[row][b] = true
                cols[col][b] = true
                boxes[row / 3][col / 3][b] = true
            }
        }
    }
    return true
}
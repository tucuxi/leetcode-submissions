func gameOfLife(board [][]int)  {
    neighbors := make([][]int, len(board))
    for i := range board {
        neighbors[i] = make([]int, len(board[i]))
    }
    for i := range board {
        for j := range board[i] {
            for p := max(i-1, 0); p <= min(i+1, len(board) - 1); p++ {
                for q := max(j-1, 0); q <= min(j+1, len(board[i]) - 1); q++ {
                    if p != i || q != j {
                       neighbors[i][j] += board[p][q]
                    }
                }
            }
        }
    }
    for i := range board {
        for j := range board[i] {
            if board[i][j] == 1 {
                if neighbors[i][j] < 2 || neighbors[i][j] > 3 {
                    board[i][j] = 0
                }
            } else {
                if neighbors[i][j] == 3 {
                    board[i][j] = 1
                } 
            }
        }
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func updateBoard(board [][]byte, click []int) [][]byte {
    r, c := click[0], click[1]
    
    if board[r][c] == 'M' {
        board[r][c] = 'X'
    } else {
        reveal(board, r, c)
    }
    
    return board
}

func visitNeighbors(board [][]byte, r, c int, f func([][]byte, int, int)) {
    for i := max(0, r-1); i <= min(len(board)-1, r+1); i++ {
        for j := max(0, c-1); j <= min(len(board[i])-1, c+1); j++ {
            if i != r || j != c {
                f(board, i, j)
            }
        }
    }
}

func reveal(board [][]byte, r, c int) {
    if board[r][c] != 'E' {
        return
    }
    
    var m MineCounter
    
    visitNeighbors(board, r, c, m.checkAndIncrement)
    if count := m.getCount(); count > 0 {
        board[r][c] = '0' + byte(count)
    } else {
        board[r][c] = 'B' 
        visitNeighbors(board, r, c, reveal)
    }
}
    
type MineCounter struct {
    count int
}

func (m *MineCounter) checkAndIncrement(board [][]byte, r, c int) {
    if board[r][c] == 'M' || board[r][c] == 'X' {
        m.count++
    }
}

func (m *MineCounter) getCount() int {
    return m.count
}
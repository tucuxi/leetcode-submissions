type node [6]int

func slidingPuzzle(board [][]int) int {
    target := node{1, 2, 3, 4, 5, 0}
    b := node{board[0][0], board[0][1], board[0][2], board[1][0], board[1][1], board[1][2]}

    seen := make(map[node]bool)
    seen[b] = false
    q := []node{b}

    move := func(i, j int) {
        newBoard := node{b[0], b[1], b[2], b[3], b[4], b[5]}
        newBoard[i], newBoard[j] = newBoard[j], newBoard[i]
        if !seen[newBoard] {
            q = append(q, newBoard)
            seen[newBoard] = true
        }
    }

    for moves := 0; len(q) > 0; moves++ {
        for range len(q) {
            b = q[0]
            q = q[1:]
            if b == target {
                return moves
            }            
            z := 0
            for ; b[z] != 0; z++ {}
            if z >= 3 {
                move(z, z - 3)
            }
            if z < 3 {
                move(z, z + 3)
            }
            if z % 3 > 0 {
                move(z, z - 1)
            }
            if z % 3 < 2 {
                move(z, z + 1)
            }
        }
    }
    return -1
}
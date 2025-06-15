const noJump = -1

func snakesAndLadders(board [][]int) int {
    n := len(board)
    target := n * n
    jumps := make(map[int]int)
    c, dc, nn := 0, 1, 1
    for r := n - 1; r >= 0; r-- {
        for i := 0; i < n; i++ {
            if board[r][c] != noJump {
                jumps[nn] = board[r][c]
            }
            c += dc
            nn++
        }
        c -= dc
        dc = -dc
    }
    // bfs
    q := []int{1}
    v := map[int]bool{1: true}
    m := 0
    for len(q) > 0 {
        for i := len(q); i > 0; i-- {
            curr := q[0]
            q = q[1:]
            if curr == target {
                return m
            }
            for j := min(curr + 6, target); j > curr; j-- {
                next, found := jumps[j]
                if !found {
                    next = j
                }
                if !v[next] {
                    v[next] = true
                    q = append(q, next)
                }
            }
        }
        m++
    }
    return -1
}
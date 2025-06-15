const maxWordLen = 10

func findWords(board [][]byte, words []string) []string {
    res := []string{}

    m := make(map[string]struct{})
    for _, w := range words {
        m[w] = struct{}{}
    }

    var dfs func(i, j int, pre []byte)
    dfs = func(i, j int, pre []byte) {
        c := board[i][j]
        if c == '$' {
            return
        }
        pre = append(pre, c)
        s := string(pre)
        if _, ok := m[s]; ok {
            res = append(res, s)
            delete(m, s)
        }
        if len(pre) == maxWordLen {
            return
        }
        board[i][j] = '$'
        if i > 0 {
            dfs(i - 1, j, pre)
        }
        if i + 1 < len(board) {
            dfs(i + 1, j, pre)
        }
        if j > 0 {
            dfs(i, j - 1, pre)
        }
        if j + 1 < len(board[i]) {
            dfs(i, j + 1, pre)
        }
        board[i][j] = c
    }

    pre := []byte{}
    for i := range board {
        for j := range board[i] {
            dfs(i, j, pre)
        }
    }
    return res
}
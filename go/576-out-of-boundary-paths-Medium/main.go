const mod = 1_000_000_007

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    memo := make([][][]int, maxMove+1)    
    for i := range memo {
        memo[i] = make([][]int, m)
        for j := range memo[i] {
            memo[i][j] = make([]int, n)
            for k := range memo[i][j] {
                memo[i][j][k] = -1
            }
        }
    }
    
    var dp func(int, int, int) int
    dp = func(move, r, c int) int {
        if r < 0 || r >= m || c < 0 || c >= n {
            return 1
        }
        if move == 0 {
            return 0
        }
        if memo[move][r][c] != -1 {
            return memo[move][r][c]
        }
        memo[move][r][c] = (dp(move-1, r-1, c) + dp(move-1, r+1, c) + dp(move-1, r, c-1) + dp(move-1, r, c+1)) % mod
        return memo[move][r][c]
    } 
    
    return dp(maxMove, startRow, startColumn)
}
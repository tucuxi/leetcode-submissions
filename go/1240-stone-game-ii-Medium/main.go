func stoneGameII(piles []int) int {
    memo := make([][][]int, 2)
    for p := range memo {
        memo[p] = make([][]int, len(piles))
        for i := range memo[p] {
            memo[p][i] = make([]int, len(piles) + 1)
            for x := range memo[p][i] {
                memo[p][i][x] = -1
            }
        }
    }
    
    var dp func(int, int, int) int
    dp = func(p, i, m int) int {
        if i == len(piles) {
            return 0
        }
        if memo[p][i][m] != -1 {
            return memo[p][i][m]
        }
        res := -1
        if p == 1 {
            res = math.MaxInt
        }
        s := 0
        for x := 1; x <= min(2 * m, len(piles) - i); x++ {
            s += piles[i + x - 1]
            if p == 0 {
                res = max(res, s + dp(1, i + x, max(m, x)))
            } else {
                res = min(res, dp(0, i + x, max(m, x)))
            }
        }
        memo[p][i][m] = res
        return res
    }
    
    return dp(0, 0, 1)
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
func minSteps(n int) int {
    if n == 1 {
        return 0
    }
    
    memo := make([][]int, n+1)
    for i := range memo {
        memo[i] = make([]int, n/2 + 1)
    }

    var dp func(int, int) int
    dp = func(currentLen, pasteLen int) int {
        if currentLen == n {
            return 0
        }
        if currentLen > n {
            return 1000
        }
        if memo[currentLen][pasteLen] != 0 {
            return memo[currentLen][pasteLen]
        }
        steps := min(dp(currentLen+pasteLen, pasteLen) + 1, dp(currentLen * 2, currentLen) + 2)
        memo[currentLen][pasteLen] = steps
        return steps
    }

    return 1 + dp(1, 1)
}
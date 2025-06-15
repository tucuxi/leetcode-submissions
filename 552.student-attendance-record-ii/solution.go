const MOD = 1_000_000_007

func checkRecord(n int) int {
    memo := make([][2][3]int, n+1)
    memo[0][0][0] = 1
    memo[0][0][1] = 1
    memo[0][0][2] = 1
    memo[0][1][0] = 1
    memo[0][1][1] = 1
    memo[0][1][2] = 1
    for i := 1; i < len(memo); i++ {
        for j := range memo[i] {
            for k := range memo[i][j] {
                memo[i][j][k] = -1
            }
        }
    }

    var dp func(int, int, int) int
    dp = func(n int, absences int, lates int) int {
        if absences >= 2 || lates >= 3 {
            return 0
        }
        if memo[n][absences][lates] != -1 {
            return memo[n][absences][lates]
        }
        res := (dp(n-1, absences, 0) + dp(n-1, absences+1, 0) + dp(n-1, absences, lates+1)) % MOD
        memo[n][absences][lates] = res
        return res
    }

    return dp(n, 0, 0)
}
const MOD = 1_000_000_007

func numberOfWays(n int, x int) int {
    p := powers(n, x)

    memo := make(map[[2]int]int)

    var dp func(int, int) int
    dp = func(n int, i int) int {
        if n < 0 {
            return 0
        }
        if n == 0 {
            return 1
        }
        if i == len(p) {
            return 0
        }
        key := [2]int{n, i}
        if ways, exists := memo[key]; exists {
            return ways
        }
        ways := (dp(n - p[i], i + 1) + dp(n, i + 1)) % MOD
        memo[key] = ways
        return ways
    }
    
    return dp(n, 0)
}

func powers(n, x int) []int {
    var res []int
    for i := 1;; i++ {
        p := 1
        for range x {
            p *= i
        }
        if p > n {
            return res
        }
        res = append(res, p)
    }
}
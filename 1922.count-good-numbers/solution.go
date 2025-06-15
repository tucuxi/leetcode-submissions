const MOD = 1_000_000_007 

func countGoodNumbers(n int64) int {
    return int(multmod(5, (n + 1) / 2) * multmod(4, n / 2) % MOD)
}

func multmod(x, y int64) int64 {
    res := int64(1)
    factor := x
    for y > 0 {
        if y % 2 == 1 {
            res = res * factor % MOD
        }
        factor = factor * factor % MOD
        y /= 2
    }
    return res
}
const mod = 1_000_000_007

func countGoodStrings(low int, high int, zero int, one int) int {
    res := 0
    ways := make([]int, high + 1)
    ways[0] = 1
    for i := 1; i <= high; i++ {
        if j := i - zero; j >=0 {
            ways[i] = ways[j]
        }
        if j := i - one; j >= 0 {
            ways[i] += ways[j]
        }
        ways[i] %= mod
        if i >= low {
            res = (res + ways[i]) % mod
        }
    }
    return res
}
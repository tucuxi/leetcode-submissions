func countNicePairs(nums []int) int {
    const mod = 1_000_000_007
    h := make(map[int]int)
    res := 0
    for _, n := range nums {
        d := rev(n) - n
        res = (res + h[d]) % mod
        h[d]++
    }
    return res
}

func rev(num int) int {
    r := 0
    for num > 0 {
        r = 10 * r + num % 10
        num /= 10
    }
    return r
}
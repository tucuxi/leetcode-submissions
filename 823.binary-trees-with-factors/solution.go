func numFactoredBinaryTrees(arr []int) int {
    const mod = 1_000_000_007
    sort.Ints(arr)
    dp := map[int]int{}
    for _, num := range arr {
        dp[num] = 1
    }
    
    for i, num := range arr {
        for _, x := range arr[:i] {
            if num % x != 0 {
                continue
            }
            if v, ok := dp[num / x]; ok {
                dp[num] += dp[x] * v
                dp[num] %= mod
            }
        } 
    }
    
    res := 0
    for _, v := range dp {
        res = (res + v) % mod
    }
    return res
}
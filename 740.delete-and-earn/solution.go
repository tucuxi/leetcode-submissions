func deleteAndEarn(nums []int) int {
    h := make(map[int]int)
    for _, n := range nums {
        h[n] += n
    }
    a := make([]int, 0, len(h))
    for k := range h {
        a = append(a, k)
    }
    sort.Ints(a)
    memo := make([]int, len(a))
    
    var dp func(int) int
    dp = func(i int) int {
        if i >= len(a) {
            return 0
        }
        if memo[i] > 0 {
            return memo[i]
        }
        if h[a[i] + 1] > 0 {
            memo[i] = max(h[a[i]] + dp(i+2), dp(i+1))
        } else {
            memo[i] = h[a[i]] + dp(i+1)
        }
        return memo[i]
    }
    
    return dp(0)
}
                
func max(a ...int) int {
    res := a[0]
    for _, v := range a {
        if v > res {
            res = v
        }
    }
    return res
}
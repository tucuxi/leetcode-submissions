func minCost(n int, cuts []int) int {
    sort.Ints(cuts)
    memo := make(map[[2]int]int)
    
    var dp func(int, int) int
    dp = func(i, j int) int {
        if j - i <= 1 {
            return 0
        }
        key := [...]int{i, j}
        if _, exists := memo[key]; !exists {
            min := math.MaxInt
            k := sort.Search(len(cuts), func(x int) bool { return cuts[x] > i })
            for k < len(cuts) && cuts[k] < j {
                if h := dp(i, cuts[k]) + dp(cuts[k], j); h < min {
                    min = h
                }
                k++
            }
            if min == math.MaxInt {
                memo[key] = 0
            } else {
                memo[key] = min + j - i
            }
        }
        return memo[key]
    }
    
    return dp(0, n)
}

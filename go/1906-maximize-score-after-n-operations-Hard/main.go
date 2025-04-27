func maxScore(nums []int) int {
    memo := make([]int, 1 << len(nums))
    
    var rec func(int, int) int
    rec = func(pair int, mask int) int {
        if pair == len(nums) / 2 || memo[mask] > 0 {
            return memo[mask]
        }
        for i := 0; i < len(nums); i++ {
            for j := i + 1; j < len(nums); j++ {
                maskPair := (1 << i) | (1 << j)
                if mask & maskPair != 0 {
                    continue
                }
                score := (pair + 1) * gcd(nums[i], nums[j]) + rec(pair + 1, mask | maskPair)
                if score > memo[mask] {
                    memo[mask] = score
                }
            }
        }
        return memo[mask]
    }
    
    return rec(0, 0)
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}
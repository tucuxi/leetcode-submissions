func minOperations(nums []int) int {
    h := make(map[int]int)
    for _, n := range nums {
        h[n]++
    }
    
    memo := make(map[int]int)
    memo[0] = 0
    memo[1] = -1
    var f func(int) int 
    f = func(n int) int {
        if n < 0 {
            return -1
        }
        if k, exists := memo[n]; exists {
            return k
        }
        if k3 := f(n-3); k3 != -1 {
            memo[n] = k3+1
        } else if k2 := f(n-2); k2 != -1 {
            memo[n] = k2+1
        } else {
            memo[n] = -1
        }
        return memo[n]
    }
    
    res := 0
    for _, c := range h {
        k := f(c)
        if k == -1 {
            return -1
        }
        res += k
    }
    return res
}
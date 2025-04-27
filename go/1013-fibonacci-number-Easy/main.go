func fib(n int) int {
    if n <= 1 {
        return n
    }
    
    memo := make([]int, n+1)
    memo[1] = 1

    var rec func(int) int
    rec = func(n int) int {
        if n <= 1 {
            return n
        }
        if memo[n] == 0 {
            memo[n] = rec(n-1) + rec(n-2)
        }
        return memo[n]
    }
    
    return rec(n)
}
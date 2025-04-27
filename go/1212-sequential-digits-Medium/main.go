func sequentialDigits(low int, high int) []int {
    var res []int
    m := length(low)
    n := length(high)
    
    for i := m; i <= n; i++ {
        for j := 1; j <= 10-i; j++ {
            n := 0
            for k := 0; k < i; k++ {
                n = 10*n + (j+k)
            }    
            if n >= low && n <= high {
                res = append(res, n)
            }
        }    
    }
    
    return res
}

func length(n int) int {
    res := 0
    for n > 0 {
        n /= 10
        res++
    }
    return res
}
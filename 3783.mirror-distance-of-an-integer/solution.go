func mirrorDistance(n int) int {
    return abs(n - reverse(n))    
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func reverse(n int) int {
    res := 0
    for ; n > 0; n /= 10 {
        res = 10*res + n%10
    }
    return res
}
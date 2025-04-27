func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n < 0 {
        n = -n
        x = 1 / x
    }
    y := 1.0
    for ; n > 1; n /= 2 {
        if n % 2 != 0 {
            y *= x
        }
        x *= x
    } 
    return x * y
}
func sumAndMultiply(n int) int64 {
    x, sum := 0, 0
    for f := 1; n > 0; n /= 10 {
        if d := n % 10; d > 0 {
            x += f * d
            sum += d
            f *= 10
        }
    }
    return int64(x) * int64(sum)
}
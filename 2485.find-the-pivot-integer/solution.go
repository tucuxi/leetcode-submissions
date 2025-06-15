func pivotInteger(n int) int {
    s := n * (n + 1) / 2
    x := int(math.Round(math.Sqrt(float64(s))))
    if x * x != s {
        return -1
    }
    return x
}
func maxProduct(n int) int {
    max1, max2 := 0, 0
    for i := n; i > 0; i /= 10 {
        d := i % 10
        if d >= max1 {
            max1, max2 = d, max1
        } else if d > max2 {
            max2 = d
        }
    }
    return max1 * max2
}
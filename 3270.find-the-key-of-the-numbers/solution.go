func generateKey(num1 int, num2 int, num3 int) int {
    res := 0
    f := 1
    for range 4 {
        res += f * min(num1%10, num2%10, num3%10)
        f *= 10
        num1 /= 10
        num2 /= 10
        num3 /= 10
    }
    return res
}
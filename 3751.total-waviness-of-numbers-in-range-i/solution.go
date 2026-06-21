func totalWaviness(num1 int, num2 int) int {
    w := 0
    for i := max(100, num1); i <= num2; i++ {
        s := strconv.Itoa(i)
        for j := 1; j < len(s) - 1; j++ {
            if s[j-1] < s[j] && s[j] > s[j+1] || s[j-1] > s[j] && s[j] < s[j+1] {
                w++
            }
        }
    }
    return w
}
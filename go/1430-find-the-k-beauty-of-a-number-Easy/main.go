func divisorSubstrings(num int, k int) int {
    res := 0
    s := strconv.Itoa(num)
    for i := k; i <= len(s); i++ {
        n, _ := strconv.Atoi(s[i-k:i])
        if n > 0 && num % n == 0 {
            res++
        }
    }
    return res
}
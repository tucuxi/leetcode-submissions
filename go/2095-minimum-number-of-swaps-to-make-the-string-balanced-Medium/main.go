func minSwaps(s string) int {
    res := 0
    balance := 0
    for _, r := range s {
        if r == '[' {
            balance++
        } else {
            balance--
            if balance < 0 {
                balance = 1
                res++
            }
        }
    }
    return res
}
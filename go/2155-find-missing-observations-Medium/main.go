func missingRolls(rolls []int, mean int, n int) []int {
    rest := mean * (len(rolls) + n) - sum(rolls)
    if rest < n || rest > 6*n {
        return nil
    }
    ans := make([]int, 0, n)
    for n > 0 {
        n--
        x := max(1, rest - 6 * n)
        rest -= x
        ans = append(ans, x)
    }
    return ans
}

func sum(nums []int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}
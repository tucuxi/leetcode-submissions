func candy(ratings []int) int {
    n := len(ratings)
    candy := make([]int, n)
    candy[0] = 1
    for i := 1; i < n; i++ {
        if ratings[i] > ratings[i - 1] {
            candy[i] = candy[i - 1] + 1
        } else {
            candy[i] = 1
        }
    }
    for i := n - 2; i >= 0; i-- {
        if ratings[i] > ratings[i + 1] && candy[i] <= candy[i + 1] {
            candy[i] = candy[i + 1] + 1
        }
    }
    res := 0
    for _, k := range candy {
        res += k
    }
    return res
}
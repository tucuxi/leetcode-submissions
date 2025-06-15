func change(amount int, coins []int) int {
    count := make([]int, amount + 1)
    count[0] = 1
    for _, c := range coins {
        for a := c; a <= amount; a++ {
            count[a] += count[a - c]
        }
    }
    return count[amount]
}
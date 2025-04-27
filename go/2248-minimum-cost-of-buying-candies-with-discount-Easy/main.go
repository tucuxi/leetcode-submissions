func minimumCost(cost []int) int {
    sort.Ints(cost)
    sum := 0
    for k, i := 0, len(cost) - 1; i >= 0; i-- {
        if k == 2 {
            k = 0
        } else {
            sum += cost[i]
            k++
        }
    }
    return sum
}
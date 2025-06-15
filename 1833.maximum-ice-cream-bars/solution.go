func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)
    k := 0
    for k < len(costs) && costs[k] <= coins {
        coins -= costs[k]
        k++
    }
    return k
}
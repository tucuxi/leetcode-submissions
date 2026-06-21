func maxIceCream(costs []int, coins int) int {
    const maxCost = 100_000
    var count [maxCost+1]int

    for _, c := range costs {
        count[c]++
    }

    res := 0
    for c, n := range count {
        if c > coins {
            break
        } 
        if n == 0 {
            continue
        }
        b := min(n, coins/c)
        res += b
        coins -= b*c
    }

    return res
}
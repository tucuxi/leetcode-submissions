func minCost(nums []int, cost []int) int64 {
    
    c := func(target int) int64 {
        var total int64
        for i := range nums {
            if ci := int64(cost[i]) * int64(nums[i] - target); ci < 0 {
                total -= ci
            } else {
                total += ci
            }
        }
        return total
    }

    min, max := cost[0], cost[0]
    for _, n := range nums {
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }
        
    n := min + sort.Search(max - min + 1, func(n int) bool {
        return c(n + min) <= c(n + min + 1)
    })
        
    return c(n)
}
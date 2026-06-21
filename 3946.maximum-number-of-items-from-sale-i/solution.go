func maximumSaleItems(items [][]int, budget int) int {
    f := make([]int, len(items))
    for i := range items {
        for j := range items {
            if j !=i && items[j][0] % items[i][0] == 0 {
                f[i]++
            }
        }
    }

    dp := make([]int, budget+1)

    for i, item := range items {
        cost := item[1]
        first := f[i] + 1

        dp2 := make([]int, budget+1)
        copy(dp2, dp)

        for b := 0; b+cost <= budget; b++ {
            dp2[b + cost] = max(dp2[b + cost], dp[b] + first)
        }

        for b := cost; b <= budget; b++ {
            dp2[b] = max(dp2[b], dp2[b - cost] + 1)
        }

        dp = dp2
    }

    res := 0

    for _, v := range dp {
        res = max(res, v)
    }

    return res
}
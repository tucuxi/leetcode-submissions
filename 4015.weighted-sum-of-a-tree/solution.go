func weightedSum(parent []int, nums []int) int64 {
    g := make([][]int, len(nums))

    for u := 1; u < len(parent); u++ {
        v := parent[u]
        g[v] = append(g[v], u)
    }

    var height func(int) int
    
    height = func(u int) int {
        maxChildHeight := 0
        for _, v := range g[u] {
            maxChildHeight = max(maxChildHeight, height(v))
        }
        return maxChildHeight+1
    }

    h := height(0)

    var dfs func(int, int) int64

    dfs = func(u, d int) int64 {
        weight := int64(h-d+1) * int64(nums[u])
        for _, v := range g[u] {
            weight += dfs(v, d+1)
        }
        return weight
    }

    return dfs(0, 1)
}
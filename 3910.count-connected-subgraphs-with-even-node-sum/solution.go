func evenSumSubgraphs(nums []int, edges [][]int) int {
    n := len(nums)
    adj := make([]uint32, n)
    for _, e := range edges {
        u := e[0]
        v := e[1]
        adj[u] |= uint32(1) << v
        adj[v] |= uint32(1) << u
    }

    res := 0
    totalMasks := uint32(1 << n)
    for m := uint32(1); m < totalMasks; m++ {
        sum := 0
        tmp := m
        for tmp != 0 {
            i := bits.TrailingZeros32(tmp)
            sum += nums[i]
            tmp &= tmp-1
        }
        if sum%2 != 0 {
            continue
        }

        start := bits.TrailingZeros32(m)
        visited := uint32(1) << start
        queue := uint32(1) << start

        for queue != 0 {
            u := bits.TrailingZeros32(queue)
            queue &= ^(uint32(1) << u)
            neighbors := adj[u] & m
            unvisited := neighbors & ^visited
            visited |= unvisited
            queue |= unvisited
        }
        if visited == m {
            res++
        }
    }

    return res
}
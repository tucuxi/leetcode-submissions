func minimumSeconds(nums []int) int {
    h := make(map[int][]int)
    for i, n := range nums {
        h[n] = append(h[n], i)
    }
    res := math.MaxInt
    for _, indexes := range h {
        dist := len(nums) - indexes[len(indexes) - 1] + indexes[0]
        for i := 1; i < len(indexes); i++ {
            dist = max(dist, indexes[i] - indexes[i-1])
        }
        res = min(res, dist / 2)
    }
    return res
}
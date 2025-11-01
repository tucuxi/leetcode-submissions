func maxDistinctElements(nums []int, k int) int {
    res := 0
    p := math.MinInt
    slices.Sort(nums)
    for _, num := range nums {
        if q := max(p, num - k); q <= num + k {
            p = q + 1
            res++
        }
    }
    return res
}
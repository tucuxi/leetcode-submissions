func frequencySort(nums []int) []int {
    h := make(map[int]int)
    for _, n := range nums {
        h[n]++
    }
    slices.SortFunc(nums, func(a, b int) int {
        if d := h[a] - h[b]; d != 0 {
            return d
        }
        return b - a
    })
    return nums
}
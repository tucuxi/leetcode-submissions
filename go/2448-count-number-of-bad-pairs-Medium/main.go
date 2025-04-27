func countBadPairs(nums []int) int64 {
    pairs := int64(len(nums) * (len(nums) - 1) / 2)
    h := make(map[int]int)
    
    for i, n := range nums {
        pairs -= int64(h[n - i])
        h[n - i]++
    }

    return pairs
}
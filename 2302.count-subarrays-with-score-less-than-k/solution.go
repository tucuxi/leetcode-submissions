func countSubarrays(nums []int, k int64) int64 {
    var (
        res, total int64
        i int
    )
    for j := range nums {
        total += int64(nums[j])
        for i <= j && total * int64(j - i + 1) >= k {
            total -= int64(nums[i])
            i++
        }
        res += int64(j - i + 1)
    }
    return res
}
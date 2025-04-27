func minimumSubarrayLength(nums []int, k int) int {
    minLen := 100
    for i := range nums {
        s := 0
        q := min(i+minLen, len(nums))
        for p := i; p < q; p++ {
            s |= nums[p]
            if s >= k {
                minLen = p-i+1
                break
            }
        }
    }
    if minLen > len(nums) {
        return -1
    }
    return minLen
}
func kLengthApart(nums []int, k int) bool {
    p := -1
    for i := range nums {
        if nums[i] == 1 {
            if p >= 0 && i - p <= k {
                return false
            }
            p = i
        }
    }
    return true
}
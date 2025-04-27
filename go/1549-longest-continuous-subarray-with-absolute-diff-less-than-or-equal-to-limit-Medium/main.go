func longestSubarray(nums []int, limit int) int {
    var (
        minDeque []int
        maxDeque []int
        maxLength = 0
    )
    for i, j := 0, 0; j < len(nums); j++ {
        for len(maxDeque) > 0 && maxDeque[len(maxDeque) - 1] < nums[j] {
            maxDeque = maxDeque[:len(maxDeque) - 1]
        }
        maxDeque = append(maxDeque, nums[j])
        for len(minDeque) > 0 && minDeque[len(minDeque) - 1] > nums[j] {
            minDeque = minDeque[:len(minDeque) - 1]
        }
        minDeque = append(minDeque, nums[j])
        for maxDeque[0] - minDeque[0] > limit {
            if maxDeque[0] == nums[i] {
                maxDeque = maxDeque[1:]
            }
            if minDeque[0] == nums[i] {
                minDeque = minDeque[1:]
            }
            i++
        }
        maxLength = max(maxLength, j - i + 1)
    }
    return maxLength
}
func hasIncreasingSubarrays(nums []int, k int) bool {
    previousSubarray, currentSubarray := 0, 0
    for i := range nums {
        if i > 0 && nums[i - 1] >= nums[i] {
            if adjacentSubarrays(previousSubarray, currentSubarray, k) {
                return true
            }
            previousSubarray, currentSubarray = currentSubarray, 0
        }
        currentSubarray++
    }
    return adjacentSubarrays(previousSubarray, currentSubarray, k)
}

func adjacentSubarrays(a, b, k int) bool {
    return max(a, b) >= 2 * k || min(a, b) >= k
}
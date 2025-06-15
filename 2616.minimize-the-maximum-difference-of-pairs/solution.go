func minimizeMax(nums []int, p int) int {
    sort.Ints(nums)
    low := 0
    high := nums[len(nums) - 1] - nums[0]
    for low < high {
        mid := (low + high) / 2
        if canFormPairs(nums, p, mid) {
            high = mid
        } else {
            low = mid + 1
        }
    }
    return high
}

func canFormPairs(nums []int, p int, m int) bool {
    count := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] - nums[i - 1] <= m {
            count++
            i++
        }
    }
    return count >= p
}
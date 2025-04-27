func countPairs(nums []int, target int) int {
    count := 0
    for i := range nums {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] < target {
                count++
            }
        }
    }
    return count
}
func minPairSum(nums []int) int {
    sort.Ints(nums)
    max := 0
    i, j := 0, len(nums) - 1
    for i < j {
        if sum := nums[i] + nums[j]; sum > max {
            max = sum
        }
        i++
        j--
    }
    return max
}
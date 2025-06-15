func reductionOperations(nums []int) int {
    sort.Ints(nums)
    res := 0
    lastIndex := len(nums) - 1
    largestNum := nums[lastIndex]
    for i := lastIndex; i >= 0; i-- {
        if nums[i] < largestNum {
            largestNum = nums[i]
            res += lastIndex - i
        }
    }
    return res
}
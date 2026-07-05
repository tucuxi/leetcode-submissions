func maxSum(nums []int, k int, mul int) int64 {
    slices.Sort(nums)

    i := len(nums)
    sum := int64(0)
    for range k {
        i--
        sum += int64(max(1, mul)) * int64(nums[i])
        mul--
    }
    return sum
}
func missingNumber(nums []int) int {
    x := len(nums) * (len(nums) + 1) / 2
    for _, n := range nums {
        x -= n
    }
    return x
}
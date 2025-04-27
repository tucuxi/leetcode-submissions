func getMaximumXor(nums []int, maximumBit int) []int {
    res := make([]int, len(nums)) 
    m := (1 << maximumBit) - 1
    x := 0
    for i := range nums {
        x ^= nums[i]
        res[len(nums) - 1 - i] = m - x
    }
    return res
}
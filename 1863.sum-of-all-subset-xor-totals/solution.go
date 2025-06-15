func subsetXORSum(nums []int) int {
    return rec(nums, 0)
}

func rec(nums []int, xor int) int {
    if len(nums) == 0 {
        return xor
    }
    sumWithNext := rec(nums[1:], xor ^ nums[0])
    sumWithoutNext := rec(nums[1:], xor)
    return sumWithNext + sumWithoutNext
}
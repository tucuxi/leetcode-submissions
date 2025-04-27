func matrixSum(nums [][]int) int {
    for i := range nums {
        sort.Ints(nums[i])
    }
    res := 0
    for j := len(nums[0]) - 1; j >= 0; j-- {
        max := 0
        for i := range nums {
            if nums[i][j] > max {
                max = nums[i][j]
            }
        }
        res += max
    }
    return res
}
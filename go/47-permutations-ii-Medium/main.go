func permuteUnique(nums []int) [][]int {
    res := [][]int{}

    var rec func(int)
    rec = func(k int) {
        if k == len(nums) {
            solution := make([]int, len(nums))
            copy(solution, nums)
            res = append(res, solution)
            return
        }
        outer: for i := k; i < len(nums); i++ {
            for j := i - 1; j >= k; j-- {
                if nums[j] == nums[i] {
                    continue outer
                }
            }
            nums[i], nums[k] = nums[k], nums[i]
            rec(k + 1)
            nums[i], nums[k] = nums[k], nums[i]
        }
    }

    rec(0)
    return res
}
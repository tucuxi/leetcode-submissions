func subsets(nums []int) [][]int {
    var res [][]int
 
    var rec func(nums []int, part []int)
    rec = func(nums []int, part []int) {
        if len(nums) == 0 {
            solution := make([]int, len(part))
            copy(solution, part)
            res = append(res, solution)
            return
        }
        rec(nums[1:], part)
        rec(nums[1:], append(part, nums[0]))
    }

    rec(nums, make([]int, 0, len(nums)))
    return res
}
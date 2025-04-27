func threeSum(nums []int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums)
    for i := range nums {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        j, k := i + 1, len(nums) - 1
        for j < k {
            if sum := nums[i] + nums[j] + nums[k]; sum == 0 {
                res = append(res, []int{nums[i], nums[j], nums[k]})
                j++
                k--
                for j < k && nums[j] == nums[j - 1] {
                    j++
                }
                for j < k && nums[k] == nums[k + 1] {
                    k--
                }
            } else if sum > 0 {
                k--
            } else {
                j++
            }
        }
    }
    return res
}
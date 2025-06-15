func permute(nums []int) [][]int {
    var (
        res [][]int
        f func(int)
    )   
    f = func(index int) {
        if index == len(nums) {
            dst := make([]int, len(nums))
            copy(dst, nums)
            res = append(res, dst)
        } else {
            for i := index; i < len(nums); i++ {
                nums[index], nums[i] = nums[i], nums[index]
                f(index + 1)
                nums[index], nums[i] = nums[i], nums[index]
            }
        }
    }
    f(0)
    return res
}
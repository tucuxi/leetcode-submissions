func incremovableSubarrayCount(nums []int) int {
    pre := make([]bool, len(nums))
    pre[0] = true
    for i := 1; i < len(nums); i++ {
        pre[i] = pre[i-1] && nums[i-1] < nums[i]
    }
    suf := make([]bool, len(nums))
    suf[len(nums)-1] = true
    for i := len(nums)-2; i >= 0; i-- {
        suf[i] = suf[i+1] && nums[i] < nums[i+1]
    }
    res := 0
    for i := 0; i < len(nums); i++ {
        if i == 0 || pre[i-1] {
            for j := i; j < len(nums); j++ {
                if j == len(nums)-1 || suf[j+1] && (i == 0 || nums[i-1] < nums[j+1]) {
                    res++
                }
            }
        }
    }
    return res
}
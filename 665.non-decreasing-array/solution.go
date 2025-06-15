func checkPossibility(nums []int) bool {
    k := 0
    p := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] < p {
            k++
            if k > 1 {
                return false
            }
            if i >= 2 && nums[i-2] > nums[i] {
                continue
            }
        }
        p = nums[i]
    }
    return true
}
func wiggleMaxLength(nums []int) int {
    if len(nums) == 1 {
        return 1
    }
    p := nums[1] - nums[0]
    l := 1
    if p != 0 {
        l++
    }
    for i := 2; i < len(nums); i++ {
        q := nums[i] - nums[i-1]
        if p <= 0 && q > 0 || p >= 0 && q < 0 {
            l++
            p = q
        }
    }
    return l
}
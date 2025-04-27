func specialArray(nums []int) int {
    slices.Sort(nums)
    for x := range len(nums) + 1 {
        k := len(nums) - x
        if (k == len(nums) || nums[k] >= x) && (k == 0 || nums[k - 1] < x) {  
           return x
        }
    }
    return -1
}
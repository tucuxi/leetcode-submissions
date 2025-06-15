func isMonotonic(nums []int) bool {
    lt, gt := 0, 0
    for i := 1; i < len(nums); i++ {
        if nums[i - 1] < nums[i] {
            lt++
        }
        if nums[i - 1] > nums[i] {
            gt++
        }
    }
    return lt == 0 || gt == 0 
}
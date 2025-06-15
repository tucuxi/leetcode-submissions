func findPeakElement(nums []int) int {
    l, r := 0, len(nums) - 1
    for l < r {
        p := l + (r - l) / 2
        if nums[p] > nums[p + 1] {
            r = p
        } else {
            l = p + 1
        }
    }
    return l
}
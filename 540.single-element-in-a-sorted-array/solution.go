func singleNonDuplicate(nums []int) int {
    l, r := 0, len(nums) - 1
    for l < r {
        m := l + (r - l) / 2 & ^1
        if nums[m] == nums[m + 1] {
            l = m + 2
        } else {
            r = m - 1
        }
    }
    return nums[l]
}
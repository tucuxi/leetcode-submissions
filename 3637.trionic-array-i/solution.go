func isTrionic(nums []int) bool {
    i := 0
    for i + 1 < len(nums) && nums[i] < nums[i + 1] {
        i++
    }
    j := i
    for j + 1 < len(nums) && nums[j] > nums[j + 1] {
        j++
    }
    k := j
    for k + 1 < len(nums) && nums[k] < nums[k + 1] {
        k++
    }
    return i > 0 && k == len(nums) - 1 && j - 1 > 0 && k - j > 0
}
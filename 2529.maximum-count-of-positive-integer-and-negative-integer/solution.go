func maximumCount(nums []int) int {
    i, _ := slices.BinarySearch(nums, 1)
    j, _ := slices.BinarySearch(nums, 0)
    return max(len(nums) - i, j)
}
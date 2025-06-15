func minOperations(nums []int, k int) int {
    slices.Sort(nums)
    i, _ := slices.BinarySearch(nums, k)
    return i
}
func partitionArray(nums []int, k int) int {
    slices.Sort(nums)
    count := 1
    i := 0
    for j := 1; j < len(nums); j++ {
        if nums[j] - nums[i] > k {
            count++
            i = j
        }
    }
    return count
}
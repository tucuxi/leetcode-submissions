func rotate(nums []int, k int)  {
    tmp := make([]int, len(nums))
    for i := range nums {
        tmp[(i + k) % len(nums)] = nums[i]
    }
    copy(nums, tmp)
}
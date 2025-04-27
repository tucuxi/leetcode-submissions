func nextPermutation(nums []int)  {
    i := len(nums) - 1
    for ; i > 0; i-- {
        if nums[i - 1] < nums[i] {
            swapIndex := i
            for j := i + 1; j < len(nums); j++ {
                if nums[j] > nums[i - 1] && nums[j] < nums[swapIndex] {
                    swapIndex = j
                }
            }
            nums[i - 1], nums[swapIndex] = nums[swapIndex], nums[i - 1]
            break
        }
    }
    sort.Ints(nums[i:])
}
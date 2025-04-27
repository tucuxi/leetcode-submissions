func searchRange(nums []int, target int) []int {
    a := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
    if a == len(nums) || nums[a] > target{
        return []int{ -1, -1 }
    }
    b := sort.Search(len(nums) - a, func(i int) bool { return nums[i + a] > target }) - 1 + a
    return []int{ a, b }
}
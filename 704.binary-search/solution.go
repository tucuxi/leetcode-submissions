func search(nums []int, target int) int {
    k := sort.Search(len(nums), func(i int) bool {
        return nums[i] >= target
    })
    if k < len(nums) && nums[k] == target {
        return k
    }
    return -1
}
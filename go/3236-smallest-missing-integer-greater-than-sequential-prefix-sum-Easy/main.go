func missingInteger(nums []int) int {
    s := nums[0]
    for i := 1; i < len(nums) && nums[i] == nums[i-1] + 1; i++ {
        s += nums[i]
    }
    sort.Ints(nums)
    k := sort.Search(len(nums), func(i int) bool {
        return nums[i] >= s
    })
    for ; k < len(nums); k++ {
        if nums[k] > s {
            break
        }
        if nums[k] == s {
            s++
        }
    }
    return s
}
func largestPerimeter(nums []int) int64 {
    sort.Ints(nums)
    var s int64
    for _, n := range nums {
        s += int64(n)
    }
    for i := len(nums) - 1; s <= 2 * int64(nums[i]); {
        s -= int64(nums[i])
        i--
        if i < 2 {
            return -1
        }
    } 
    return s
}
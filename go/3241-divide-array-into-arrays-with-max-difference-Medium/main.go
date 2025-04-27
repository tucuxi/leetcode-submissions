func divideArray(nums []int, k int) [][]int {
    var res [][]int
    
    sort.Ints(nums)
    for i := 0; i < len(nums); i += 3 {
        if nums[i+2] - nums[i] > k {
            return nil
        }
        res = append(res, nums[i:i+3])
    }
    return res
}
func moveZeroes(nums []int)  {
    k := 0
    for _, n := range nums {
        if n != 0 {
            nums[k] = n
            k++
        }
    }
    for k < len(nums) {
        nums[k] = 0
        k++
    }
}
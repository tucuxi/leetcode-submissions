func transformArray(nums []int) []int {
    evens := 0
    for _, num := range nums {
        if num % 2 == 0 {
            evens++
        }
    }
    for i := range nums {
        if evens > 0 {
            nums[i] = 0
            evens--
        } else {
            nums[i] = 1
        }
    }
    return nums
}
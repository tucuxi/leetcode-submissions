func minimumCost(nums []int) int {
    min1, min2 := math.MaxInt, math.MaxInt
    for _, n := range nums[1:] {
        if n < min1 {
            min1, min2 = n, min1
        } else if n < min2 {
            min2 = n
        }
    }
    return nums[0] + min1 + min2
}
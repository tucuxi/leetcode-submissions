func numberOfArithmeticSlices(nums []int) int {
    res := 0
    h := make(map[[2]int]int)   
    for i := range nums {
        for j := i-1; j >= 0; j-- {
            d := nums[i] - nums[j]
            n := h[[2]int{j, d}]
            h[[2]int{i, d}] += n+1
            res += n
        }
    }
    return res
}
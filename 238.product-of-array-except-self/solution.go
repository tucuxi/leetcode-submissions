func productExceptSelf(nums []int) []int {
    n := len(nums)
    answer := make([]int, n)
    answer[0] = 1
    for i := 1; i < n; i++ {
        answer[i] = answer[i - 1] * nums[i - 1] 
    }
    r := 1
    for i := n - 1; i >= 0; i-- {
        answer[i] *= r
        r *= nums[i]
    }
    return answer
}
func isGood(nums []int) bool {
    n := len(nums) - 1
    f := make([]int, len(nums))
    for _, num := range nums {
        if num < 1 || num > n {
            return false
        }
        f[num]++
    }
    for i := 1; i < n; i++ {
        if f[i] != 1 {
            return false
        }
    }
    return f[n] == 2
}
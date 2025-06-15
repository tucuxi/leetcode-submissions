func checkSubarraySum(nums []int, k int) bool {
    pre := map[int]int{0: 0}
    sum := 0
    for i, n := range nums {
        sum = (sum + n) % k
        if p, ok := pre[sum]; !ok {
            pre[sum] = i + 1
        } else if p < i {
            return true
        }
    }
    return false
}
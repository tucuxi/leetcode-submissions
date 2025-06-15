func subarraySum(nums []int, k int) int {
    frq := map[int]int{}
    sum := 0
    res := 0
    for _, n := range nums {
        sum += n
        if sum == k {
            res++
        }
        res += frq[sum - k]
        frq[sum]++
    }
    return res
}
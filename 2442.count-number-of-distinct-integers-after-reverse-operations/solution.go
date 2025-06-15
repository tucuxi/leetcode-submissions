func countDistinctIntegers(nums []int) int {
    distinct := make(map[int]bool)
    for _, n := range nums {
        distinct[n] = true
        r := 0
        for ; n > 0; n /=10 {
            r = 10 * r + n % 10
        }
        distinct[r] = true
    }
    return len(distinct)
}
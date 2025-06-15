func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
    var (
        res int64
        pre int
        cnt = map[int]int{0: 1}
    )

    for _, num := range nums {
        if num % modulo == k {
            pre++
        }
        res += int64(cnt[(pre - k + modulo) % modulo])
        cnt[pre % modulo]++
    }

    return res
}
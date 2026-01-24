const MOD = 1_000_000_007

func specialTriplets(nums []int) int {
    prefix := make(map[int]int)
    suffix := make(map[int]int)

    for _, n := range nums {
        suffix[n]++
    }

    var res int64

    for _, n := range nums {
        suffix[n]--
        res = (res + int64(prefix[2*n]) * int64(suffix[2*n])) % MOD
        prefix[n]++
    }

    return int(res)
}
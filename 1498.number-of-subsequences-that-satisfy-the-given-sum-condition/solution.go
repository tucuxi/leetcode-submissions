const mod = 1_000_000_007

func numSubseq(nums []int, target int) int {
    e := exp(len(nums))
    sort.Ints(nums)
    res := 0
    for p, q := 0, len(nums) - 1; p <=q; {
        if nums[p] + nums[q] > target {
            q--
        } else {
            res = (res + e[q - p]) % mod
            p++
        }
    } 
    return res
}

func exp(n int) []int {
    e := make([]int, n)
    e[0] = 1
    for i := 1; i < n; i++ {
        e[i] = 2 * e[i - 1] % mod
    }
    return e
}
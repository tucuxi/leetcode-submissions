const mod = 1_000_000_007

func rangeSum(nums []int, n int, left int, right int) int {
    var sums []int

    for i := range nums {
        s := 0
        for j := i; j < n; j++ {
            s += nums[j]
            sums = append(sums, s)
        }
    }

    slices.Sort(sums)

    res := 0
    for _, s := range sums[left - 1:right] {
        res = (res + s) % mod
    }

    return res
}
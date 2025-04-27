func maximumSum(nums []int) int {
    res := -1
    h := make(map[int]int)
    for _, n := range nums {
        ds := digitsSum(n)
        currentMax := h[ds]
        if currentMax > 0 {
            res = max(res, currentMax + n)
        }
        h[ds] = max(currentMax, n)
    }
    return res
}

func digitsSum(n int) int {
    ds := 0
    for ; n > 0; n /= 10 {
        ds += n % 10
    }
    return ds
}
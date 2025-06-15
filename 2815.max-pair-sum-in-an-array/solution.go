func maxSum(nums []int) int {
    var h [10]int
    maxSum := -1
    for _, n := range nums {
        d := maxDigit(n)
        if h[d] > 0 && h[d] + n > maxSum {
            maxSum = h[d] + n
        }
        if n > h[d] {
            h[d] = n
        }
    }
    return maxSum
}

func maxDigit(a int) int {
    res := 0
    for ; a > 0; a /= 10 {
        d := a % 10
        if d > res {
            res = d
        }
    }
    return res
}
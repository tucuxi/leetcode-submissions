func maximizeExpressionOfThree(nums []int) int {
    a := math.MinInt
    b := math.MinInt
    c := math.MaxInt
    for _, num := range nums {
        if num > a {
            b = a
            a = num
        } else if num > b {
            b = num
        }
        if num < c {
            c = num
        }
    }
    return a + b - c
}
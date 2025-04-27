func thirdMax(nums []int) int {
    max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt
    for _, n := range nums {
        if n > max1 {
            max1, max2, max3 = n, max1, max2
        } else if n == max1 {
            continue
        } else if n > max2 {
            max2, max3 = n, max2
        } else if n == max2 {
            continue
        } else if n > max3 {
            max3 = n
        }
    }
    if max3 == math.MinInt {
        return max1
    }
    return max3
}
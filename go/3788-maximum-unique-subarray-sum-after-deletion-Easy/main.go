func maxSum(nums []int) int {
    var h [101]bool
    taken := false
    largestNegative := math.MinInt
    sum := 0

    for _, num := range nums {
        if num >= 0 {
            if !h[num] {
                sum += num
                h[num] = true
                taken = true
            }
        } else {
            largestNegative = max(largestNegative, num)
        }
    }

    if taken {
        return sum
    }
    return largestNegative
}
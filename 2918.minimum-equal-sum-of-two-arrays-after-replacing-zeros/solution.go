func minSum(nums1 []int, nums2 []int) int64 {
    minSum1, canIncrease1 := check(nums1)
    minSum2, canIncrease2 := check(nums2)
    switch {
        case minSum1 == minSum2:
            return minSum1
        case minSum1 < minSum2 && canIncrease1:
            return minSum2
        case minSum2 < minSum1 && canIncrease2:
            return minSum1
        default:
            return -1
    }
}

func check(nums []int) (int64, bool) {
    var (
        sum    int64
        zeroes bool
    )
    for _, num := range nums {
        if num == 0 {
            zeroes = true
            sum++
        } else {
            sum += int64(num)
        }
    }
    return sum, zeroes
}
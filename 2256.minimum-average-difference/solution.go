func minimumAverageDifference(nums []int) int {
    l := 0
    r := 0
    for _, n := range nums {
        r += n
    }
    minDiff, minDiffIdx := math.MaxInt, 0
    for i := range nums {
        l += nums[i]
        r -= nums[i]
        rDiff := 0
        if r != 0 {
            rDiff = r / (len(nums) - i - 1)
        }
        if diff := abs(l / (i + 1) - rDiff); diff < minDiff {
            minDiff = diff
            minDiffIdx = i
        }
    }
    return minDiffIdx
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
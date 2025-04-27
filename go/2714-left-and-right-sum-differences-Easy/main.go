func leftRigthDifference(nums []int) []int {
    l, r := 0, 0
    for _, n := range nums {
        r += n
    }
    answer := make([]int, len(nums))
    for i := range nums {
        r -= nums[i]
        answer[i] = abs(l - r)
        l += nums[i]
    }
    return answer
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
func findDuplicates(nums []int) []int {
    var res []int
    for _, n := range nums {
        i := abs(n) - 1
        if nums[i] < 0 {
            res = append(res, i + 1)
        }
        nums[i] = -nums[i]
    }
    return res
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
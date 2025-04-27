func maximizeGreatness(nums []int) int {
    sort.Ints(nums)
    res := 0
    for i, j := 0, 0; j < len(nums); j++ {
        if nums[i] < nums[j] {
            res++
            i++
        }
    }
    return res
}
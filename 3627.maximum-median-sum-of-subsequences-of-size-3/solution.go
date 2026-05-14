func maximumMedianSum(nums []int) int64 {
    var res int64

    slices.Sort(nums)

    i, j := 0, len(nums)-2

    for i < j {
        res += int64(nums[j])
        i++
        j -= 2
    }

    return res
}
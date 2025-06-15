func summaryRanges(nums []int) []string {
    var res []string
    i, j := 0, 1
    for i + j < len(nums) {
        if nums[i] + j != nums[i + j] {
            res = append(res, makeRange(nums[i], j))
            i += j
            j = 0
        }
        j++
    }
    if i < len(nums) {
        res = append(res, makeRange(nums[i], j))
    }
    return res
}

func makeRange(start, elems int) string {
    if elems == 1 {
        return fmt.Sprintf("%d", start)
    }
    return fmt.Sprintf("%d->%d", start, start + elems - 1)
}
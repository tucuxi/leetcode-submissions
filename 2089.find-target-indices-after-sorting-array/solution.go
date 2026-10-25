func targetIndices(nums []int, target int) []int {
    var res []int
    
    slices.Sort(nums)

    for i, num := range nums {
        if num == target {
            res = append(res, i)
        }
    }

    return res
}
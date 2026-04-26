func distance(nums []int) []int64 {
    h := make(map[int][]int)

    for i, num := range nums {
        h[num] = append(h[num], i)
    }

    res := make([]int64, len(nums))

    for _, g := range h {
        total := int64(0)
        for _, index := range g {
            total += int64(index)
        }
        prefix := int64(0)
        for i, index := range g {
            res[index] = total - 2*prefix + int64(index)*int64(2*i-len(g))
            prefix += int64(index)
        }
    }
    return res
}
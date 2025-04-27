func minimumOperations(nums []int) int {
    h := make(map[int]int)
    for _, num := range nums {
        h[num]++
    }

    nonDistincts := 0
    for _, freq := range h {
        if freq > 1 {
            nonDistincts++
        }
    }

    ops := 0
    for nonDistincts > 0 {
        for range min(3, len(nums)) {
            h[nums[0]]--
            if h[nums[0]] == 1 {
                nonDistincts--
            }
            nums = nums[1:]
        }
        ops++
    }

    return ops
}
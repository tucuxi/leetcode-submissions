func arrayChange(nums []int, operations [][]int) []int {
    h := make(map[int]int)
    for i, n := range nums {
        h[n] = i
    }

    for _, o := range operations {
        a, b := o[0], o[1]
        i := h[a]
        nums[i] = b
        delete(h, a)
        h[b] = i
    }

    return nums
}
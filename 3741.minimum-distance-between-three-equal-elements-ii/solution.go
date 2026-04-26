func minimumDistance(nums []int) int {
    m := len(nums) + 1
    first := make([]int, m)
    second := make([]int, m)
    res := math.MaxInt

    for i := range m {
        first[i] = -1
        second[i] = -1
    }
    for i, num := range nums {
        if first[num] == -1 {
            first[num] = i
        } else if second[num] == -1 {
            second[num] = i
        } else {
            res = min(res, 2 * (i-first[num]))
            first[num] = second[num]
            second[num] = i
        }
    }
    if res == math.MaxInt {
        return -1
    }
    return res
}
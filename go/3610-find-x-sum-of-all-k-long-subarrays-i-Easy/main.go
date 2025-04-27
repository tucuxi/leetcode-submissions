func findXSum(nums []int, k int, x int) []int {
    n := len(nums)
    h := make(map[int]int)
    for _, n := range nums[:k] {
        h[n]++
    }
    answer := make([]int, 0, n - k + 1)
    answer = append(answer, xSum(h, x))
    for i := k; i < n; i++ {
        h[nums[i - k]]--
        h[nums[i]]++
        answer = append(answer, xSum(h, x))
    }
    return answer
}

func xSum(h map[int]int, x int) int {
    var f [][2]int
    for n, c := range h {
        f = append(f, [2]int{n, c})
    }
    slices.SortFunc(f, func(a, b [2]int) int {
        if a[1] == b[1] {
            return b[0] - a[0]
        }
        return b[1] - a[1]
    })
    res := 0
    for i := min(x, len(f)) - 1; i >= 0; i-- {
        res += f[i][0] * f[i][1]
    }
    return res
}
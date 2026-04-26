func minMirrorPairDistance(nums []int) int {
    h := make(map[int]int)
    res := math.MaxInt

    for j, num := range nums {
        if i, ok := h[num]; ok {
            res = min(res, j-i)
        }
        h[reverse(num)] = j
    }

    if res == math.MaxInt {
        return -1
    }

    return res

}

func reverse(num int) int {
    res := 0
    for ; num > 0; num /= 10 {
        res = 10 * res + num % 10
    }
    return res
}

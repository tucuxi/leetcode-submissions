func decrypt(code []int, k int) []int {
    res := make([]int, len(code))
    sign := sgn(k)
    for i := range code {
        sum := 0
        l := i
        for j := sign * k; j > 0; j-- {
            l += sign
            if l == len(code) {
                l = 0
            }
            if l < 0 {
                l += len(code)
            }
            sum += code[l]
        }
        res[i] = sum
    }
    return res
}

func sgn(a int) int {
    if a < 0 {
        return -1
    }
    if a > 0 {
        return 1
    }
    return 0
}
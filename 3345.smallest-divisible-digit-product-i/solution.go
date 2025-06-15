func smallestNumber(n int, t int) int {
    k := n
    for prod(digits(k)) % t != 0 {
        k++
    }
    return k
}

func digits(n int) []int {
    var res []int
    for ; n > 0; n /= 10 {
        res = append(res, n % 10)
    }
    return res
}

func prod(digits []int) int {
    res := 1
    for _, d := range digits {
        res *= d
    }
    return res
}
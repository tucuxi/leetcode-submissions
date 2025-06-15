func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    res := make([]int, n)
    pa := make([]bool, n + 1)
    pb := make([]bool, n + 1)
    c := 0
    for i := range n {
        if a := A[i]; !pa[a] {
            pa[a] = true
            if pb[a] {
                c++
            }
        }
        if b := B[i]; !pb[b] {
            pb[b] = true
            if pa[b] {
                c++
            }
        }
        res[i] = c
    }
    return res
}
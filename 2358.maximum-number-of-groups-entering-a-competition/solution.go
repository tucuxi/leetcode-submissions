func maximumGroups(grades []int) int {
    k := 0
    for i := 0; i + k < len(grades); i += k {
        k++
    }
    return k
}
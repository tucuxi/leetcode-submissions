func occurrencesOfElement(nums []int, queries []int, x int) []int {
    occ := []int{-1}
    for i, n := range nums {
        if n == x {
            occ = append(occ, i)
        }
    }
    answer := make([]int, 0, len(queries))
    for _, q := range queries {
        if q >= len(occ) {
            answer = append(answer, -1)
        } else {
            answer = append(answer, occ[q])
        }
    }
    return answer
}
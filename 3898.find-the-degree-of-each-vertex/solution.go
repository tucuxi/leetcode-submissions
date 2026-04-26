func findDegrees(matrix [][]int) []int {
    ans := make([]int, len(matrix))
    for i, r := range matrix {
        for _, c := range r {
            ans[i] += c
        }
    }
    return ans
}
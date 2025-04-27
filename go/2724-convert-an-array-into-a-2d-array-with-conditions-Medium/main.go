func findMatrix(nums []int) [][]int {
    var res [][]int
    var h [201]int
    
    for _, n := range nums {
        i := h[n]
        h[n] = i+1
        if i == len(res) {
            res = append(res, []int{n})
        } else {
            res[i] = append(res[i], n)
        }
    }
    
    return res
}
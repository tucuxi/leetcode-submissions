func getSneakyNumbers(nums []int) []int {
    var res []int
    h := make([]bool, len(nums) - 2)
    for _, n := range nums {
        if h[n] {
            res = append(res, n)
        } else {
            h[n] = true
        }
    } 
    return res
}
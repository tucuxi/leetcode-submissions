func findDisappearedNumbers(nums []int) []int {
    h := make([]bool, len(nums) + 1)
    for _, n := range nums {
        h[n] = true
    }
    var res []int
    for i := 1; i < len(h); i++ {
        if !h[i] {
            res = append(res, i)
        }
    }
    return res
}
func numIdenticalPairs(nums []int) int {
    res := 0
    var h [101]int
    for _, n := range nums {
        res += h[n]
        h[n]++
    } 
    return res
}
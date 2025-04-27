func isPossibleToSplit(nums []int) bool {
    h := make(map[int]int)  
    for _, n := range nums {
        h[n]++
        if h[n] > 2 {
            return false
        }
    }
    return true
}
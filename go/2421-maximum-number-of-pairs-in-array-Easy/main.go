func numberOfPairs(nums []int) []int {
    h := map[int]int{}
    for _, n := range nums {
        h[n]++
    }
    pairs, leftovers := 0, 0
    for _, f := range h {
        pairs += f / 2
        leftovers += f % 2
    }
    return []int{pairs, leftovers}
}
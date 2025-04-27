func findLucky(arr []int) int {
    freq := map[int]int{}
    for _, v := range arr {
        freq[v]++
    }
    maxLucky := -1
    for _, v := range arr {
        if v == freq[v] {
            if v > maxLucky {
                maxLucky = v
            }
        }
    }
    return maxLucky
}
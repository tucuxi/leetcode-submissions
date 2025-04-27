func uniqueOccurrences(arr []int) bool {
    h := make(map[int]int)
    for _, a := range arr {
        h[a]++
    }
    s := make(map[int]bool)
    for _, v := range h {
        if s[v] {
            return false
        }
        s[v] = true
    }
    return true
}
func firstUniqChar(s string) int {
    var h [26]int
    for _, r := range s {
        h[r - 'a']++
    }
    for i, r := range s {
        if h[r - 'a'] == 1 {
            return i
        }
    }
    return -1
}
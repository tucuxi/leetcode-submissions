func buddyStrings(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }
    if s == goal {
        h := map[byte]bool{}
        for i := range s {
            h[s[i]] = true
        }
        return len(h) < len(s) 
    }
    d := []int{}
    for i := range s {
        if s[i] != goal[i] {
            d = append(d, i)
        }
    }
    if len(d) != 2 {
        return false
    }
    return s[d[0]] == goal[d[1]] && s[d[1]] == goal[d[0]]
}
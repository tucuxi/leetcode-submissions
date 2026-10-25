func maxDepth(s string) int {
    depth, maxDepth := 0, 0
    for _, r := range s {
        switch r {
        case '(':
            depth++
            maxDepth = max(depth, maxDepth)
        case ')':
            depth--
        }
    }
    return maxDepth
}
func removeStars(s string) string {
    stack := make([]rune, 0, len(s))
    for _, r := range s {
        if r == '*' {
            stack = stack[:len(stack) - 1]
        } else {
            stack = append(stack, r)
        }
    }
    return string(stack)
}
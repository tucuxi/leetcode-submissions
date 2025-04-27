func minLength(s string) int {
    stack := make([]rune, 0, len(s))
    for _, b := range s {
        stack = append(stack, b)
        for len(stack) >= 2 {
            top := string(stack[len(stack) - 2:])
            if top != "AB" && top != "CD" {
                break
            }
            stack = stack[:len(stack) - 2]
        }
    }
    return len(stack)
}
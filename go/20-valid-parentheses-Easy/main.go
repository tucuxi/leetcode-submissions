func isValid(s string) bool {
    paren := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }
    stack := []rune{}
    for _, r := range s {
        if closing, ok := paren[r]; ok {
            stack = append(stack, closing)
        } else if len(stack) > 0 && stack[len(stack) - 1] == r {
            stack = stack[:len(stack) - 1]
        } else {
            return false
        }
    }
    return len(stack) == 0
}
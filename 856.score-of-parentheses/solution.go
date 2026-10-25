func scoreOfParentheses(s string) int {
    stack := []int{0}
    for _, r := range s {
        if r == '(' {
            stack = append(stack, 0)
        } else {
            v := stack[len(stack) - 1]
            w := stack[len(stack) - 2]
            stack = append(stack[:len(stack) - 2], w + max(2*v, 1))
        }
    }
    return stack[0]
}
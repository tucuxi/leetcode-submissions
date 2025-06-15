func validateStackSequences(pushed []int, popped []int) bool {
    stack := make([]int, 0, len(pushed))
    j := 0
    for _, n := range pushed {
        stack = append(stack, n)
        for j < len(popped) && len(stack) > 0 && stack[len(stack) - 1] == popped[j] {
            stack = stack[:len(stack) - 1]
            j++
        }
    }
    return len(stack) == 0
}
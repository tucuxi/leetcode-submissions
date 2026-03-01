func numSteps(s string) int {
    steps := 0
    carry := 0

    for i := len(s) - 1; i > 0; i-- {
        digit := int(s[i] - '0') + carry
        if digit == 1 {
            steps++
            carry = 1
        }
        steps++
    }

    return steps + carry
}
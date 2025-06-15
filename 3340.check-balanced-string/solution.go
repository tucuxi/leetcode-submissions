func isBalanced(num string) bool {
    b := 0
    s := 1
    for _, n := range num {
        b += s * int(n - '0')
        s = -s
    }
    return b == 0
}
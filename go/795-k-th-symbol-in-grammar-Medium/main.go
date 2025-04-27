func kthGrammar(n int, k int) int {
    if n == 1 {
        return 0
    }
    return kthGrammar(n - 1, (k + 1) / 2) ^ ((k + 1) & 1)
}
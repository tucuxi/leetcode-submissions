func finalValueAfterOperations(operations []string) int {
    res := 0
    for _, o := range operations {
        res += eval(o)
    }
    return res
}

func eval(s string) int {
    if s[1] == '-' {
        return -1
    }
    return 1
}
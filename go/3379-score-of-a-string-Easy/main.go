func scoreOfString(s string) int {
    sc := 0
    for i := 1; i < len(s); i++ {
        sc += abs(int(s[i]) - int(s[i-1]))
    }
    return sc
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
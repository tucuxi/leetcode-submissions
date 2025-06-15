func hasSpecialSubstring(s string, k int) bool {
    previous := rune(0)
    run := 0
    for _, r := range s {
        if r == previous {
            run++
            continue
        }
        if run == k {
            break
        }
        previous = r
        run = 1
    }
    return run == k
}
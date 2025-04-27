func doesValidArrayExist(derived []int) bool {
    p := 0
    for _, d := range derived {
        p ^= d
    }
    return p == 0
}

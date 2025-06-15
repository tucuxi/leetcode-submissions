func checkStrings(s1 string, s2 string) bool {
    h := make(map[byte]int)
    for i := 0; i < len(s1); i += 2 {
        update(h, s1[i], 1)
        update(h, s2[i], -1)
    }
    if len(h) > 0 {
        return false
    }
    for i := 1; i < len(s1); i += 2 {
        update(h, s1[i], 1)
        update(h, s2[i], -1)
    }
    return len(h) == 0
}

func update(h map[byte]int, b byte, v int) {
    h[b] += v
    if h[b] == 0 {
        delete(h, b)
    }
}
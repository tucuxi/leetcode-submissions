func canBeEqual(s1 string, s2 string) bool {
    b1 := []byte(s1)
    b2 := []byte(s2)
    if b1[0] != b2[0] {
        b2[0], b2[2] = b2[2], b2[0]
    }
    if b1[1] != b2[1] {
        b2[1], b2[3] = b2[3], b2[1]
    }
    return b1[0] == b2[0] && b1[1] == b2[1] && b1[2] == b2[2] && b1[3] == b2[3]
}
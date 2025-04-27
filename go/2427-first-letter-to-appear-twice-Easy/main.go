func repeatedCharacter(s string) byte {
    seen := map[byte]bool{}
    for i := range s {
        c := s[i]
        if seen[c] {
            return c
        }
        seen[c] = true
    }
    panic("no letter appears twice")
}
func passwordStrength(password string) int {
    h := make(map[rune]bool)
    strength := 0

    for _, ch := range password {
        if !h[ch] {
            h[ch] = true
            if ch >= 'a' && ch <= 'z' {
                strength++
            } else if ch >= 'A' && ch <= 'Z' {
                strength += 2
            } else if ch >= '0' && ch <= '9' {
                strength += 3
            } else if strings.IndexRune("!@#$", ch) >= 0 {
                strength += 5
            }
        }
    }
    return strength
}
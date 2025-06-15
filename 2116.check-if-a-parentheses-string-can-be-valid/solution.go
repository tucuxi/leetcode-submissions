func canBeValid(s string, locked string) bool {
    if len(s) % 2 == 1 {
        return false
    }

    var openParen, unlocked []int

    for i := range s {
        if locked[i] == '0' {
            unlocked = append(unlocked, i)
        } else if s[i] == '(' {
            openParen = append(openParen, i)
        } else if len(openParen) > 0 {
            openParen = openParen[:len(openParen) - 1]
        } else if len(unlocked) > 0 {
            unlocked = unlocked[:len(unlocked) - 1]
        } else {
            return false
        }
    }

    for len(openParen) > 0 && len(unlocked) > 0 && openParen[len(openParen) - 1] < unlocked[len(unlocked) -1] {
        openParen = openParen[:len(openParen) - 1]
        unlocked = unlocked[:len(unlocked) - 1]
    }

    return len(openParen) == 0
}
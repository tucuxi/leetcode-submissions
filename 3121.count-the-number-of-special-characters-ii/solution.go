func numberOfSpecialChars(word string) int {
    var lower, upper, exclude [26]bool

    for _, c := range word {
        if c < 'a' {
            k := c - 'A'
            upper[k] = true
        } else {
            k := c - 'a'
            lower[k] = true
            exclude[k] = exclude[k] || upper[k]
        }
    }

    res := 0

    for k := range lower {
        if lower[k] && upper[k] && !exclude[k] {
            res++
        }
    }

    return res
}
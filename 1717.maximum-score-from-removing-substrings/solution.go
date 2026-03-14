func maximumGain(s string, x int, y int) int {
    var p1, p2 int
    var t string
    
    if (x >= y) {
        t, p1 = removeSubstring(s, 'a', 'b', x)
        _, p2 = removeSubstring(t, 'b', 'a', y)
    } else {
        t, p1 = removeSubstring(s, 'b', 'a', y)
        _, p2 = removeSubstring(t, 'a', 'b', x)
    }
    return p1 + p2
}

func removeSubstring(s string, a, b rune, p int) (string, int) {
    var t []rune
    points := 0

    for _, ch := range s {
        if ch == b && len(t) > 0 && t[len(t) - 1] == a {
            t = t[:len(t) - 1]
            points += p
        } else {
            t = append(t, ch)
        }
    }

    return string(t), points
}
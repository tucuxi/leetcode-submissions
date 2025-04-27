func furthestDistanceFromOrigin(moves string) int {
    var d, u int
    for _, m := range moves {
        switch m {
        case 'L':
            d++
        case 'R':
            d--
        case '_':
            u++
        }
    }
    if d < 0 {
        d = -d
    }
    return u + d
}
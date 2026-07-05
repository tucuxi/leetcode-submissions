func maxDistance(moves string) int {
    x := 0
    y := 0
    w := 0

    for _, m := range moves {
        switch m {
            case 'U':
                y++
            case 'D':
                y--
            case 'L':
                x--
            case 'R':
                x++
            case '_':
                w++ 
        }
    }
    return abs(x) + abs(y) + w
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
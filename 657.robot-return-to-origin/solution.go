func judgeCircle(moves string) bool {
    r := 0
    c := 0
    for _, m := range moves {
        switch m {
        case 'R': c++
        case 'L': c--
        case 'U': r--
        case 'D': r++
        }
    }
    return r == 0 && c == 0
}
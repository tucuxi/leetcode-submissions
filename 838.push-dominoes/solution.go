func pushDominoes(dominoes string) string {
    n := len(dominoes)
    f := make([]int, n)

    force := 0
    for i := range dominoes {
        switch dominoes[i] {
            case 'L': force = 0
            case 'R': force = n
            default: force = max(force - 1, 0)
        }
        f[i] += force
    }

    force = 0
    for i := n - 1; i >= 0; i-- {
        switch dominoes[i] {
            case 'L': force = n
            case 'R': force = 0
            default: force = max(force - 1, 0)
        }
        f[i] -= force
    }

    var sb strings.Builder
    for _, force := range f {
        var d rune
        switch {
            case force < 0: d = 'L'
            case force > 0: d = 'R'
            default: d = '.' 
        }
        sb.WriteRune(d)
    }

    return sb.String()
}
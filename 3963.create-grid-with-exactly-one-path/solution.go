func createGrid(m int, n int) []string {
    var (
        res []string
        sb  strings.Builder
    )

    sb.WriteByte('.')

    for range n-1 {
        sb.WriteByte('#')
    }

    for range m-1 {
        res = append(res, sb.String())
    }

    sb.Reset()
    
    for range n {
        sb.WriteByte('.')
    }

    res = append(res, sb.String())
    return res
}
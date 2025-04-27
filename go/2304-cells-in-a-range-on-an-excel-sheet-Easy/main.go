func cellsInRange(s string) []string {
    res := []string{}
    var col1, col2 byte
    var row1, row2 int
    fmt.Sscanf(s, "%c%d:%c%d", &col1, &row1, &col2, &row2)
    for y := col1; y <= col2; y++ {
        for x := row1; x <= row2; x++ {
            res = append(res, fmt.Sprintf("%c%d", y, x))
        }
    }
    return res
}
func addSpaces(s string, spaces []int) string {
    var b strings.Builder
    spaceIndex := 0

    for i := range s {
        if spaceIndex < len(spaces) && i == spaces[spaceIndex] {
            b.WriteByte(' ')
            spaceIndex++
        }
        b.WriteByte(s[i])
    }
    return b.String()
}
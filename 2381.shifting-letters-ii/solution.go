func shiftingLetters(s string, shifts [][]int) string {
    changes := make([]int, len(s) + 1)
    for _, shift := range shifts {
        start, end, d := shift[0], shift[1], 2 * shift[2] - 1
        changes[start] += d
        changes[end + 1] -= d
    }

    var sb strings.Builder
    offset := 0
    for i, ch := range s {
        offset += changes[i]
        ch = rune(((((int(ch - 'a') + offset) % 26) + 26) % 26) + 'a')
        sb.WriteRune(ch)
    }
    
    return sb.String()
}
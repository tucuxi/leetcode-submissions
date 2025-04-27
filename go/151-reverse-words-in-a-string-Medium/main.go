func reverseWords(s string) string {
    words := []string{}
    var w strings.Builder
    for i := range s {
        if s[i] != ' ' {
            w.WriteByte(s[i])
        }
        if (s[i] == ' ' || i == len(s) - 1) && w.Len() > 0 {
            words = append([]string{w.String()}, words...)
            w.Reset()
        }
    }
    return strings.Join(words, " ")
}
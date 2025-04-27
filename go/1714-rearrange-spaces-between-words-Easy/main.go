func reorderSpaces(text string) string {
    spaces := 0
    words := []int{}
    inword := false
    for i := range text {
        if text[i] == ' ' {
            spaces++
            inword = false
        } else if !inword {
            inword = true
            words = append(words, i)
        }
    }
    inter := 0
    if len(words) > 1 {
        inter = spaces / (len(words) - 1)
    }
    var sb strings.Builder
    for i := 0; i < len(words); i++ {
        for j := words[i]; j < len(text) && text[j] != ' '; j++ {
            sb.WriteByte(text[j])
        }
        if i < len(words) - 1 {
            for j := inter; j > 0; j-- {
                sb.WriteByte(' ')
            }
        } else {
            for j := spaces - inter * (len(words) - 1); j > 0; j-- {
                sb.WriteByte(' ')
            }
        }
    }
    return sb.String()    
}
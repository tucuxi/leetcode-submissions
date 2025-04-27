func maxRepeating(sequence string, word string) int {
    for k := len(sequence) / len(word); k > 0; k-- {
        var sb strings.Builder
        for i := 0; i < k; i++ {
            sb.WriteString(word)
        }
        if strings.Contains(sequence, sb.String()) {
            return k
        }
    }
    return 0
}
func findOcurrences(text string, first string, second string) []string {
    var result []string
    words := strings.Fields(text)
    for i := 1; i < len(words) - 1; i++ {
        if (words[i] == second && words[i - 1] == first) {
            result = append(result, words[i + 1])
        }
    }
    return result
}
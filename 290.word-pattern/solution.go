func wordPattern(pattern string, s string) bool {
    patternMap := map[byte]string{}
    wordMap := map[string]byte{}
    words := strings.Split(s, " ")
    if len(words) != len(pattern) {
        return false
    }
    for i := range pattern {
        if _, ok := patternMap[pattern[i]]; !ok {
            patternMap[pattern[i]] = words[i]
        }
        if _, ok := wordMap[words[i]]; !ok {
            wordMap[words[i]] = pattern[i]
        }
        if patternMap[pattern[i]] != words[i] || wordMap[words[i]] != pattern[i] {
            return false
        }
    }
    return true
}
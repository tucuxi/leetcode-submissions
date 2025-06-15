func isPrefixOfWord(sentence string, searchWord string) int {
    skipping := false
    index := 1
    matchedCharacters := 0
    for _, r := range sentence {
        if r == ' ' {
            skipping = false
            index++
            matchedCharacters = 0
            continue
        }
        if skipping {
            continue
        }
        if byte(r) != searchWord[matchedCharacters] {
            skipping = true
            continue
        }
        matchedCharacters++
        if matchedCharacters == len(searchWord) {
            return index
        }
    }
    return -1
}
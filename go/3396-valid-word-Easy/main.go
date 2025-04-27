var vocals = "aeiouAEIOU"

func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }
    var v, c bool
    for _, ch := range word {
        if unicode.In(ch, unicode.Digit) {
            continue
        }
        if !unicode.In(ch, unicode.Letter) {
            return false
        }
        if strings.ContainsRune(vocals, ch) {
            v = true
        } else {
            c = true
        }
    }
    return v && c
}
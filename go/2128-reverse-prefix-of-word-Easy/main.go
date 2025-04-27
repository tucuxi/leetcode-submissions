func reversePrefix(word string, ch byte) string {
    i := 0
    for {
        if i == len(word) {
            return word
        }
        if word[i] == ch {
            break
        }
        i++
    }

    var r []byte
    for j := i; j >= 0; j-- {
        r = append(r, word[j])
    }    
    r = append(r, word[i+1:]...)
    return string(r)
}
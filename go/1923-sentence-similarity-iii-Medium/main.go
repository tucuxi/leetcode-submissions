func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    words1 := strings.Split(sentence1, " ")
    words2 := strings.Split(sentence2, " ")

    l := 0
    for l < len(words1) && l < len(words2) && words1[l] == words2[l] {
        l++
    }

    r1 := len(words1) - 1
    r2 := len(words2) - 1
    for r1 >= 0 && r2 >= 0 && words1[r1] == words2[r2] {
        r1--
        r2--
    }

    return r1 < l || r2 < l
}

/*

a b c
d e a a b c

*/
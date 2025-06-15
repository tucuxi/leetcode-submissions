func countOfSubstrings(word string, k int) int64 {
    return atLeastK(word, k) - atLeastK(word, k + 1)
}

const VOWELS = "aeiou"

func atLeastK(word string, k int) int64 {
    res := int64(0)
    vowels := make(map[byte]int)
    consonants := 0
    i := 0

    for j := range word {
        chJ := word[j]
        if strings.IndexByte(VOWELS, chJ) >= 0 {
            vowels[chJ]++
        } else {
            consonants++
        }
        for ; len(vowels) == len(VOWELS) && consonants >= k; i++ {
            res += int64(len(word) - j)
            chI := word[i]
            if strings.IndexByte(VOWELS, chI) >= 0 {
                if vowels[chI] == 1 {
                    delete(vowels, chI)
                } else {
                    vowels[chI]--
                }
            } else {
                consonants--
            }
        }
    }
    return res
}
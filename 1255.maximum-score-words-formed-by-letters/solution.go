func maxScoreWords(words []string, letters []byte, score []int) int {
    counts := histogram(letters)
    maxScore := 0

    var comb func(int, int)
    comb = func(index, currentScore int) {
        if index == len(words) {
            maxScore = max(currentScore, maxScore)
            return
        }
        comb(index + 1, currentScore)
        canUseWord := true
        wordScore := 0
        for _, r := range words[index] {
            c := byte(r) - 'a'
            wordScore += score[c]
            counts[c]--
            if counts[c] < 0 {
                canUseWord = false
            }
        }
        if canUseWord {
            comb(index + 1, currentScore + wordScore)
        }
        for _, r := range words[index] {
            c := byte(r) - 'a'
            counts[c]++
        }
    }

    comb(0, 0)
    return maxScore
}

func histogram(letters []byte) []int {
    res := make([]int, 26)
    for _, r := range letters {
        c := byte(r) - 'a'
        res[c]++
    }
    return res
}
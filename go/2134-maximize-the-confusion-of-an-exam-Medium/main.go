func maxConsecutiveAnswers(answerKey string, k int) int {
    res := make(chan int)
    defer close(res)

    solve := func(consecutiveAnswer byte) {
        i := 0
        r := k
        for j := range answerKey {
            if answerKey[j] != consecutiveAnswer {
                r--
            }
            if r < 0 {
                if answerKey[i] != consecutiveAnswer {
                    r++
                }
                i++
            }
        }
        res <- len(answerKey) - i
    }

    go solve('T')
    go solve('F')

    return max(<-res, <-res)
}
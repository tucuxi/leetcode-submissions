func bagOfTokensScore(tokens []int, power int) int {
    sort.Ints(tokens)
    score := 0
    max := 0
    for i, j := 0, len(tokens) - 1; i <= j; {
        if power >= tokens[i] {
            power -= tokens[i]
            score++
            if score > max {
                max = score
            }
            i++
        } else if score > 0 {
            power += tokens[j]
            score--
            j--
        } else {
            break
        }
    }
    return max
}
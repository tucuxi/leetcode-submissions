func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    return sum(firstWord) + sum(secondWord) == sum(targetWord)
}

func sum(s string) int {
    res := 0
    for i := range s {
        res = res * 10 + int(s[i] - 'a')
    }
    return res
}
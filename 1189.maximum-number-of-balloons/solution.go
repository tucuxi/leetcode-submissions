func maxNumberOfBalloons(text string) int {
    var c [26]int

    for _, ch := range text {
        c[ch-'a']++
    }

    return min(c[1], c[0], c[11] / 2, c[14] / 2, c[13])
}
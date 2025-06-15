func getHint(secret string, guess string) string {
    hsecret := map[byte]int{}
    hguess := map[byte]int{}
    bulls := 0
    cows := 0
    for i := range guess {
        if guess[i] == secret[i] {
            bulls++
        } else {
            hsecret[secret[i]]++
            hguess[guess[i]]++
        }
    }
    for ch, count := range hguess {
        cows += min(count, hsecret[ch])
    }
    return fmt.Sprintf("%dA%dB", bulls, cows)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func equalFrequency(word string) bool {
    freq, maxFreq := map[rune]int{}, 0
    for _, c := range word {
        freq[c]++
        if freq[c] > maxFreq {
            maxFreq = freq[c]
        }
    }
    maxFreqCount := 0
    minFreq := maxFreq
    for _, f := range freq {
        if f == maxFreq {
            maxFreqCount++
        } else if f < minFreq {
            minFreq = f
        }
    }
    return len(freq) == 1 ||
            len(freq) == len(word) ||
            maxFreqCount == 1 && maxFreq - minFreq == 1 ||
            len(freq) - maxFreqCount == 1 && minFreq == 1
}
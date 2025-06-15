func maxDifference(s string) int {
    var count [26]int

    for _, r := range s {
        count[r - 'a']++
    }

    maxOddFrequency := 0
    minEvenFrequency := math.MaxInt

    for _, c := range count {
        if c % 2 != 0 {
            maxOddFrequency = max(maxOddFrequency, c)
        } else if c > 0 {
            minEvenFrequency = min(minEvenFrequency, c)
        }
    }

    if minEvenFrequency == math.MaxInt {
        return maxOddFrequency
    }
    return maxOddFrequency - minEvenFrequency
}
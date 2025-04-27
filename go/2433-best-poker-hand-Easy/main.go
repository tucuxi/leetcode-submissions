func bestHand(ranks []int, suits []byte) string {
    if flush(ranks, suits) {
        return "Flush"
    }
    if threeOfKind(ranks, suits) {
        return "Three of a Kind"
    }
    if pair(ranks, suits) {
        return "Pair"
    }
    return "High Card"
}

func flush(_ []int, suits []byte) bool {
    for i := 1; i < len(suits); i++ {
        if suits[i] != suits[0] {
            return false
        }
    }
    return true
}

func threeOfKind(ranks []int, _ []byte) bool {
    freq := [14]int{}
    for _, r := range ranks {
        freq[r]++
    }
    for _, f := range freq {
        if f >= 3 {
            return true
        }
    }
    return false
}

func pair(ranks []int, _ []byte) bool {
    freq := [14]int{}
    for _, r := range ranks {
        freq[r]++
    }
    for _, f := range freq {
        if f == 2 {
            return true
        }
    }
    return false
}
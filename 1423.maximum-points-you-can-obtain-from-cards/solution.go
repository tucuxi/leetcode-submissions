func maxScore(cardPoints []int, k int) int {
    w := 0
    for i := 0; i < k; i++ {
        w += cardPoints[i]
    }
    a, l := w, len(cardPoints)
    for i := 1; i <= k; i++ {
        w += cardPoints[l-i] - cardPoints[k-i]
        if w > a {
            a = w
        }
    }
    return a
}
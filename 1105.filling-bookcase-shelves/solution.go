func minHeightShelves(books [][]int, shelfWidth int) int {
    memo := make([][]int, len(books))
    for i := range memo {
        memo[i] = make([]int, shelfWidth+1)
    }

    var f func(int, int, int) int
    f = func(i, remainingWidth, maxHeight int) int {
        if i == len(books) {
            return maxHeight
        }
        if memo[i][remainingWidth] > 0 {
            return memo[i][remainingWidth]
        }

        h := maxHeight + f(i + 1, shelfWidth - books[i][0], books[i][1])
        if books[i][0] <= remainingWidth {
            h = min(h, f(i + 1, remainingWidth - books[i][0], max(maxHeight, books[i][1])))
        }
        memo[i][remainingWidth] = h
        return h
    }

    return f(0, shelfWidth, 0)
}
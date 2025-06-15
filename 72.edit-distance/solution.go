type transformation struct {
    cost int
    d1 int
    d2 int
}

type operation int

const (
    right operation = iota
    insert
    remove
    replace
)

var transform = []transformation{
    { 0, -1, -1 },   // right
    { 1,  0, -1 },   // insert
    { 1, -1,  0 },   // remove
    { 1, -1, -1 },   // replace
}

func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m + 1)
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }
    for i := 0; i <= m; i++ {
        for j := 0; j <= n; j++ {
            if i == 0 && j == 0 {
                continue
            }
            dp[i][j] = math.MaxInt
            if i > 0 && j > 0 && word1[i - 1] == word2[j - 1] {
                updateIfCheaper(dp, i, j, transform[right])
            }
            for _, op := range []operation{insert, remove, replace} {
                updateIfCheaper(dp, i, j, transform[op])
            }
        }
    }
    return dp[m][n]
}

func updateIfCheaper(dp [][]int, i, j int, t transformation) {
    if i >= -t.d1 && j >= -t.d2 {
        if alt := dp[i + t.d1][j + t.d2] + t.cost; alt < dp[i][j] {
            dp[i][j] = alt
        }
    }
}
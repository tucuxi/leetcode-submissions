func winningPlayerCount(n int, pick [][]int) int {
    var count [10][11]int
    for _, p := range pick {
        count[p[0]][p[1]]++
    }
    res := 0
    for player := range n {
        for color := range count[player] {
            if count[player][color] > player {
                res++
                break
            }
        }
    }
    return res
}
func findChampion(n int, edges [][]int) int {
    teams := make(map[int]bool)
    for i := range n {
        teams[i] = true
    }
    for _, edge := range edges {
        delete(teams, edge[1])
    }
    if len(teams) == 1 {
        for winner := range teams {
            return winner
        }
    }
    return -1
}
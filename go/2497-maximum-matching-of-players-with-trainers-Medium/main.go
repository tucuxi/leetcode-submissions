func matchPlayersAndTrainers(players []int, trainers []int) int {
    slices.Sort(players)
    slices.Sort(trainers)
    i, j := 0, 0
    for i < len(players) && j < len(trainers) {
        if players[i] <= trainers[j] {
            i++
        }
        j++
    }
    return i
}
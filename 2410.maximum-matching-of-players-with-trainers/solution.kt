class Solution {
    fun matchPlayersAndTrainers(players: IntArray, trainers: IntArray): Int {
        players.sort()
        return trainers.sorted().fold(0) { i, tc ->
            if (i < players.size && players[i] <= tc) {
                i + 1
            } else {
                i
            }
        }
    }
}
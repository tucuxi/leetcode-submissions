class Solution {
    fun unhappyFriends(n: Int, preferences: Array<IntArray>, pairs: Array<IntArray>): Int {        
        val pairedFriends = pairs
            .flatMap { (a, b) -> listOf(a to b, b to a) }
            .sortedBy { (a, _) -> a }
            .map { (_, b) -> b }
            .toIntArray()

        return pairedFriends.withIndex().count { (x, y) ->
            preferences[x].takeWhile { it != y }.any { u ->
                preferences[u].indexOf(x) < preferences[u].indexOf(pairedFriends[u])
            }
        }
    }
}
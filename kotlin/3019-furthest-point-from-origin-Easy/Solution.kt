class Solution {
    fun furthestDistanceFromOrigin(moves: String): Int {
        val (d, u) = moves.fold(Pair(0, 0)) { (d, u), ch ->
            when (ch) {
                'L' -> Pair(d + 1, u)
                'R' -> Pair(d - 1, u)
                '_' -> Pair(d, u + 1)
                else -> error("invalid input")
            }
        }
        return abs(d) + u
    }
}
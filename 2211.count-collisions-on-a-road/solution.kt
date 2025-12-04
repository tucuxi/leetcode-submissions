class Solution {
    fun countCollisions(directions: String): Int {
        return directions
            .dropWhile { it == 'L' }
            .dropLastWhile { it == 'R' }
            .count { it != 'S' }
    }
}
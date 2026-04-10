class Solution {
    fun judgeCircle(moves: String): Boolean {
        val (r, c) = moves
            .fold(0 to 0) { (r, c), m ->
                when (m) { 
                    'R' -> r to c + 1
                    'L' -> r to c - 1
                    'U' -> r - 1 to c
                    'D' -> r + 1 to c
                    else -> throw IllegalArgumentException()
                }
            }
    
        return r == 0 && c == 0
    }
}
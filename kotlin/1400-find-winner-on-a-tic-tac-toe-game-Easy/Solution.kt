class Solution {
    fun tictactoe(moves: Array<IntArray>): String {
        val horizontal = IntArray(3)
        val vertical = IntArray(3)
        var diagonal1 = 0
        var diagonal2 = 0  
        for (i in moves.indices) {
            val w = weight(i)
            val x = moves[i][0]
            val y = moves[i][1]
            horizontal[y] += w
            vertical[x] += w
            if (x == y)
                diagonal1 += w
            if (x + y == 2)
                diagonal2 += w
            val t = target(i)
            if (horizontal[y] == t
                    || vertical[x] == t
                    || diagonal1 == t
                    || diagonal2 == t)
                return winner(i)
        }
        if (moves.size < 9)
            return "Pending"
        return "Draw"
    }
    
    private fun weight(moveIndex: Int) = (moveIndex % 2) * 2 - 1
    private fun target(moveIndex: Int) = (moveIndex % 2) * 6 - 3
    private fun winner(moveIndex: Int) = if (moveIndex % 2 == 0) "A" else "B"
}
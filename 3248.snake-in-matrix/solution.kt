class Solution {

    fun finalPositionOfSnake(n: Int, commands: List<String>): Int {
        val step = mapOf("UP" to -n, "RIGHT" to 1, "DOWN" to n, "LEFT" to -1)
        return commands.mapNotNull { step[it] }.sum()
    }
}
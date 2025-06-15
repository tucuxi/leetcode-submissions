class Solution {
    fun countPoints(rings: String): Int {
        val rods = Array(10) { mutableSetOf<Char>() }
        rings.windowed(2, 2).forEach {
            rods[it[1] - '0'].add(it[0])
        }
        return rods.count { it.size == 3 }
    }
}
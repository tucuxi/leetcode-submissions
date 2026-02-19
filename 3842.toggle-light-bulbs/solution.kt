class Solution {
    fun toggleLightBulbs(bulbs: List<Int>): List<Int> {
        val b = BooleanArray(101)
        bulbs.forEach { b[it] = !b[it] }
        return (1..b.lastIndex).filter { b[it] }
    }
}
class Solution {
    fun findLucky(arr: IntArray): Int {
        return arr
            .asSequence()
            .groupingBy { it }
            .eachCount()
            .filter { it.key == it.value }
            .maxByOrNull { it.key }
            ?.key
        ?: -1 
    }
}
class Solution {
    fun sortByBits(arr: IntArray): IntArray {
        val sorted = arr.sortedWith(
            compareBy(
                { it.countOneBits() },
                { it }
            )
        )
        return sorted.toIntArray()
    }
}
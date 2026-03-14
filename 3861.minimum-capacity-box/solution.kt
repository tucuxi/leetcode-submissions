class Solution {
    fun minimumIndex(capacity: IntArray, itemSize: Int): Int {
        return capacity.withIndex()
            .filter { it.value >= itemSize }
            .minByOrNull { it.value }
            ?.index ?: -1
    }
}
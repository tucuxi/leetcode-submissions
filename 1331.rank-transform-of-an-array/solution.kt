class Solution {
    fun arrayRankTransform(arr: IntArray): IntArray {
        val uniqueValues = arr.toSet().sorted()
        return arr.map { uniqueValues.binarySearch(it) + 1 }.toIntArray()
    }
}
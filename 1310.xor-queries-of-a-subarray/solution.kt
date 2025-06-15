class Solution {
    fun xorQueries(arr: IntArray, queries: Array<IntArray>): IntArray {
        val prefix = arr.runningFold(0) { acc, elem -> acc xor elem }
        return queries.map { (left, right) -> prefix[left] xor prefix[right + 1] }.toIntArray()
    }
}
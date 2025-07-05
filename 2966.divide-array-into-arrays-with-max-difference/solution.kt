class Solution {
    fun divideArray(nums: IntArray, k: Int): Array<IntArray> {
        val res = nums
            .sorted()
            .chunked(3)
            .map { it.toIntArray() }
            .toTypedArray()
        return if (res.all { it[2] - it[0] <= k }) {
            res
        } else {
            emptyArray()
        }
    }
}
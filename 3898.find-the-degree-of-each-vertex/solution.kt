class Solution {
    fun findDegrees(matrix: Array<IntArray>): IntArray {
        return matrix.map { it.sum() }.toIntArray()
    }
}
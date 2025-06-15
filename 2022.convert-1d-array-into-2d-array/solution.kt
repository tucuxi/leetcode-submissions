class Solution {
    fun construct2DArray(original: IntArray, m: Int, n: Int): Array<IntArray> {
        return if (original.size != m * n) emptyArray()
        else Array(m) { original.sliceArray(it * n until (it + 1) * n) }
    }
}
class Solution {
    fun kWeakestRows(mat: Array<IntArray>, k: Int): IntArray {
        val sum = mat.map { row -> row.count { it == 1 } }
        val kWeakest = (0 until mat.size).sortedBy { sum[it] }.take(k)
        return kWeakest.toIntArray()
    }
}
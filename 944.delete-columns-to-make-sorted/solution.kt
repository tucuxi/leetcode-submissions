class Solution {
    fun minDeletionSize(strs: Array<String>): Int {
        val deletedColumns = BooleanArray(strs.first().length)
        for (i in 1 until strs.size) {
            for (j in deletedColumns.indices) {
                deletedColumns[j] = deletedColumns[j] or (strs[i][j] < strs[i - 1][j])
            }
        }
        return deletedColumns.count { it }
    }
}
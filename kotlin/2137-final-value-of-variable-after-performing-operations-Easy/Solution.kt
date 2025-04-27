class Solution {
    fun finalValueAfterOperations(operations: Array<String>): Int =
        operations.size - 2 * operations.count { it[1] == '-' }
}
class Solution {
    fun finalValueAfterOperations(operations: Array<String>): Int {
        return operations.sumOf { eval(it) }
    }

    inline fun eval(s: String) = if (s[1] == '-') -1 else 1
}
class Solution {
    fun minOperations(s: String): Int =
        s.filter { ch -> ch > 'a' }.minOrNull()
            ?.let { smallest ->'z' - smallest + 1 }
            ?: 0
}
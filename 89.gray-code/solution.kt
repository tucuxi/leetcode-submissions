class Solution {
    fun grayCode(n: Int): List<Int> {
        return IntArray(1 shl n) { it xor (it shr 1) }.asList()
    }
}
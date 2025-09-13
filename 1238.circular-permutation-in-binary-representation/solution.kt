class Solution {
    fun circularPermutation(n: Int, start: Int): List<Int> =
        IntArray(1 shl n) { it xor (it shr 1) xor start }.asList()
}
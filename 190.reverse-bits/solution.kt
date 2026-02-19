class Solution {
    fun reverseBits(n: Int): Int {
        var cur = n
        var res = 0

        repeat(32) {
            res = (res shl 1) or (cur and 1)
            cur = cur ushr 1
        }
        return res
    }
}
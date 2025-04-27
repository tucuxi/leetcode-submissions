class Solution {
    fun findComplement(num: Int): Int {
        var p = num
        var q = 0
        while (p > 0) {
            q = (q shl 1) or 1
            p = p shr 1
        }
        return num xor q 
    }
}
class Solution {
    fun isOneBitCharacter(bits: IntArray): Boolean {
        var i = 0
        var one = true
        while (i < bits.size) {
            one = (bits[i] == 0)
            if (!one) i++
            i++
        }
        return one
    }
}
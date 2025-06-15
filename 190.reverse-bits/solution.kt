class Solution {
    // you need treat n as an unsigned value
    fun reverseBits(n:Int):Int {
        var res = 0
        var mask = 1
        while (mask != 0) {
            res = res shl 1
            if (n and mask != 0)
                res = res or 1
            mask = mask shl 1
        }
        return res
    }
}
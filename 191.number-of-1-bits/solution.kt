class Solution {
    // you need treat n as an unsigned value
    fun hammingWeight(n:Int):Int {
        var count = 0
        var nprime = n
        while (nprime != 0) {
            nprime = nprime and nprime - 1
            count++
        }
        return count
    }
}
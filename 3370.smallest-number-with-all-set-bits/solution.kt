class Solution {
    fun smallestNumber(n: Int): Int {
        var b = 0
        var n = n
        while (n > 0) {
            b++
            n /= 2
        }
        return (1 shl b) - 1
    }
}
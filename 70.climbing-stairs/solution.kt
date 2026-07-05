class Solution {
    fun climbStairs(n: Int): Int {
        var a = 1
        var b = 1

        repeat(n - 1) {
            a = b.also { b += a }
        }
        return b
    }
}
class Solution {
    fun maximizeExpressionOfThree(nums: IntArray): Int {
        var a = Int.MIN_VALUE
        var b = Int.MIN_VALUE
        var c = Int.MAX_VALUE
        nums.forEach { num ->
            if (num > a) {
                b = a
                a = num
            } else {
                b = maxOf(b, num)
            }
            c = minOf(c, num)
        }
        return a + b - c
    }
}
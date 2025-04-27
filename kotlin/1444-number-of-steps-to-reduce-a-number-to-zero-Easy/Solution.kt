class Solution {
    fun numberOfSteps(num: Int): Int =
        rec(num)
        
    private fun rec(num: Int, acc: Int = 0): Int =
        if (num == 0) {
            acc
        } else {
            val n = if (num % 2 == 0) {
                num / 2
            } else {
                num - 1
            }
            rec(n, acc + 1)
        }
}
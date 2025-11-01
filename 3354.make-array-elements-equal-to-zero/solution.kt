class Solution {
    fun countValidSelections(nums: IntArray): Int {
        val total = nums.sum()
        var res = 0
        var prefix = 0

        nums.forEach { n ->
            prefix += n
            if (n == 0) {
                res += when(total - 2 * prefix) {
                    0     -> 2
                    -1, 1 -> 1
                    else  -> 0
                }
            }
        }

        return res
    }
}
class Solution {
    fun minOperations(nums: IntArray): Int {
        val count = mutableMapOf<Int, Int>()
        var nonDistincts = 0

        nums.forEach { num ->
            count[num] = 1 + count.getOrElse(num) { 0 }
            if (count[num] == 2) {
                nonDistincts++
            }
        }

        var res = 0
        var i = 0

        while (nonDistincts > 0) {
            repeat(minOf(3, nums.size - i)) {
                val num = nums[i]
                val numCount = count.getValue(num) - 1
                if (numCount == 1) {
                    nonDistincts--
                }
                count[num] = numCount 
                i++
            }
            res++
        }
        return res
    }
}
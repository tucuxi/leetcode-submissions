class Solution {
    fun sortColors(nums: IntArray): Unit {
        val count = IntArray(3)
        nums.forEach { count[it]++ }
        var k = 0
        count.forEachIndexed { i, n ->
            repeat(n) { nums[k++] = i }
        }
    }
}
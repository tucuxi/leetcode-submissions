class Solution {
    fun sortArrayByParityII(nums: IntArray): IntArray {
        val (even, odd) = nums.partition{ it % 2 == 0 }
        return IntArray(nums.size) { if (it % 2 == 0) even[it / 2] else odd[it / 2] }
    }
}
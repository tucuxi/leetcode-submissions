class Solution {
    fun singleNumber(nums: IntArray): IntArray {
        val x = nums.reduce { acc, num -> acc xor num }
        val r = x and -x
        val a = nums.filter { it and r != 0 }.reduce { acc, num -> acc xor num }
        return intArrayOf(a, a xor x)   
    }
}
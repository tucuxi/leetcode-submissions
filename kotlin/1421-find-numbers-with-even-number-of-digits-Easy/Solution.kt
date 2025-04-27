class Solution {
    fun findNumbers(nums: IntArray): Int {
       return nums.filter { it.toString().length % 2 == 0 }.count()
    }
}
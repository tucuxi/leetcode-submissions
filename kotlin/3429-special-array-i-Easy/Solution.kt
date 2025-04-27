class Solution {
    fun isArraySpecial(nums: IntArray) = nums.asIterable().zipWithNext().all { (a, b) -> a and 1 != b and 1 } 
}
class Solution {
    fun constructTransformedArray(nums: IntArray): IntArray =    
        nums.mapIndexed { index, value -> nums[(index + value % nums.size + nums.size) % nums.size] }.toIntArray()
}
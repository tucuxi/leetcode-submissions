class Solution {
    fun getMinDistance(nums: IntArray, target: Int, start: Int): Int {
        var k = 0
        while (true) {
            if (start - k >= 0 && nums[start - k] == target) {
                break
            }
            if (start + k < nums.size && nums[start + k] == target) {
                break
            }
            k++
        }
        return k
    }
}
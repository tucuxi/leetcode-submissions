class Solution {
    fun hasSameDigits(s: String): Boolean {
        var nums = s.map { it - '0' }
        while (nums.size > 2) {
            nums = nums.zipWithNext { a, b -> (a + b) % 10 }
        }
        return nums[0] == nums[1]
    }
}
class Solution {
    fun minOperations(nums: IntArray, k: Int): Int {
        return nums.fold(0) { rest, num -> (num + rest) % k }
    }
}
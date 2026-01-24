class Solution {
    fun minOperations(nums: IntArray, target: IntArray): Int {
        val valuesToUpdate = nums
            .zip(target)
            .filterNot { (a, b) -> a == b }
            .map { (a, _) -> a }
            .toSet()
        return valuesToUpdate.size
    }
}
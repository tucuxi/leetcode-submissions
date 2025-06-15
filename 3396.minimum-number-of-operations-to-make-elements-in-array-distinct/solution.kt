class Solution {
    fun minimumOperations(nums: IntArray): Int {
        var arr = nums.toList()
        var res = 0
        while (1 < arr.groupingBy { it }.eachCount().maxOfOrNull { it.value } ?: 0) {
            arr = arr.drop(3)
            res++
        }
        return res
    }
}
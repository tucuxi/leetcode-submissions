class Solution {
    fun minOperations(nums: IntArray): Int {
        val stack = mutableListOf<Int>()
        var res = 0
        nums.forEach { a ->
            while (stack.isNotEmpty() && stack.last() > a) {
                stack.removeLast()
            }
            if (a > 0 && (stack.isEmpty() || stack.last() < a)) {
                res++
                stack.add(a)
            }
        }
        return res
    }
}
class Solution {
    fun getSneakyNumbers(nums: IntArray): IntArray {
        val n = nums.size - 2
        val h = BooleanArray(n)
        val res = mutableListOf<Int>()
        nums.forEach { num ->
            if (h[num]) {
                res.add(num)
            } else {
                h[num] = true
            }
        }
        return res.toIntArray()
    }
}
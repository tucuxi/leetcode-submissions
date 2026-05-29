class Solution {
    fun minElement(nums: IntArray): Int {
        return nums.minOf {
            var s = 0
            var r = it

            while (r > 0) {
                s += r % 10
                r /= 10
            }

            s
        }
    }
}
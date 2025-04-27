class Solution {
    fun majorityElement(nums: IntArray): Int {
        class Acc(val major: Int, val count: Int)
        val res = nums.fold(Acc(0, 0)) { acc, num ->
            if (acc.count == 0) {
                Acc(num, 1)
            } else if (acc.major == num) {
                Acc(acc.major, acc.count + 1)
            } else {
                Acc(acc.major, acc.count - 1)
            }
        }
        return res.major
    }
}
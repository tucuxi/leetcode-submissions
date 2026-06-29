class Solution {
    fun maximumElementAfterDecrementingAndRearranging(arr: IntArray): Int {
        arr.sort()
        var res = 0
        for (a in arr) {
            if (a > res) {
                res++
            }
        }
        return res
    }
}
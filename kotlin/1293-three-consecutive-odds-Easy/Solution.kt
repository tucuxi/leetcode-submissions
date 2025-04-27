class Solution {
    fun threeConsecutiveOdds(arr: IntArray): Boolean {
        var oddSequenceCount = 0
        for (elem in arr) {
            if (elem % 2 == 0) {
                oddSequenceCount = 0
            } else {
                if (++oddSequenceCount == 3) {
                    return true
                }
            }
        }
        return false
    }
}
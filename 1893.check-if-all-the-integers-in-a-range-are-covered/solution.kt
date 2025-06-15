class Solution {
    fun isCovered(ranges: Array<IntArray>, left: Int, right: Int): Boolean {
        for (x in left..right) {
            if (ranges.none { (start, end) -> x in start..end }) {
                return false
            }
        }
        return true
    }
}
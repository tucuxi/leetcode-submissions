class Solution {
    fun titleToNumber(columnTitle: String): Int {
        return columnTitle.fold(0) { acc, ch -> 26 * acc + (ch - 'A').toInt() + 1 }
    }
}
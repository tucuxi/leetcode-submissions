class Solution {
    fun countNegatives(grid: Array<IntArray>) =
        grid.fold(0) { acc, elem -> acc + elem.count { it < 0 } }
}
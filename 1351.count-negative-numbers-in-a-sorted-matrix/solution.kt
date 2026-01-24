class Solution {
    fun countNegatives(grid: Array<IntArray>): Int {
        return grid.sumOf { row -> row.size - row.strictBinarySearch { it < 0 } }
    }
}

// Unlike binarySearch from the standard library, this reports the
// index of the first match or the insert position.
fun IntArray.strictBinarySearch(predicate: (Int) -> Boolean): Int {
    var l = 0
    var r = size

    while (l < r) {
        val m = l + (r - l) / 2
        if (!predicate(this[m])) {
            l = m + 1
        } else {
            r = m
        }
    }
    return r
}
class Solution {
    fun minimumBoxes(apple: IntArray, capacity: IntArray): Int {
        var apples = apple.sum()
        capacity.sortDescending()
        val taken = capacity.takeWhile { box -> (apples > 0).also { apples -= box } }
        return taken.size
    }
}
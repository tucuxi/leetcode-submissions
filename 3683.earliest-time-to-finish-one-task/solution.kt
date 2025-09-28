class Solution {
    fun earliestTime(tasks: Array<IntArray>): Int {
        return tasks.minOf { (s, t) -> s + t }
    }
}
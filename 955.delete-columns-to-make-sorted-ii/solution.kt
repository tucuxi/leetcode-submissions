class Solution {
    fun minDeletionSize(strs: Array<String>): Int {
        var result = List(strs.size) { "" }
        return strs.first().indices.count { i ->
            val tentative = (result zip strs).map { (a, b) -> a + b[i] }
            val conflict = tentative.zipWithNext().any { (a, b) -> a > b }
            if (!conflict) result = tentative
            conflict 
        }
    }
}
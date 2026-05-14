class Solution {
    fun rotateString(s: String, goal: String): Boolean {
        return goal.length == s.length && goal in s + s
    }
}
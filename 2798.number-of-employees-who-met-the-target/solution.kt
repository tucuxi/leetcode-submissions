class Solution {
    fun numberOfEmployeesWhoMetTarget(hours: IntArray, target: Int) =
        hours.count { it >= target }
}
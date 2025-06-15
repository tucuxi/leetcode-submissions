class Solution {
    fun busyStudent(startTime: IntArray, endTime: IntArray, queryTime: Int) =
        startTime.zip(endTime).count { queryTime in it.first .. it.second }
}
class Solution {
    fun countStudents(students: IntArray, sandwiches: IntArray): Int {
        val pref = Array(2) { index -> students.count { it == index } }
        for (s in sandwiches) {
            if (pref[s] == 0) break
            pref[s]--
        }
        return pref.sum()
    }
}
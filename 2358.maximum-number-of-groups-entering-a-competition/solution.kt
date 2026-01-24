class Solution {
    fun maximumGroups(grades: IntArray): Int {
        var k = 0
        var i = 0
        while (i + k < grades.size) {
            k++
            i += k
        }
        return k
    }
}
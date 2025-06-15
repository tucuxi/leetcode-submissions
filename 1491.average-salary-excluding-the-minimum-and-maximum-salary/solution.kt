class Solution {
    fun average(salary: IntArray): Double {
        var min = 1_000_000
        var max = 0
        var sum = 0L
        for (s in salary) {
            min = minOf(min, s)
            max = maxOf(max, s)
            sum += s
        }
        return (sum - min - max).toDouble() / (salary.size - 2)
    }
}
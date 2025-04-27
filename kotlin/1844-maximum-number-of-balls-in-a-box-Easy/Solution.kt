class Solution {
    fun countBalls(lowLimit: Int, highLimit: Int): Int {
        val digitSums = (lowLimit..highLimit).map { digitSum(it) }
        return digitSums.groupingBy { it }.eachCount().values.max() ?: 0
    }
    
    fun digitSum(n: Int): Int {
        var num = n
        var sum = 0
        while (num > 0) {
            sum += num % 10
            num /= 10
        }
        return sum
    }
}
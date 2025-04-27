class Solution {
    fun checkPerfectNumber(num: Int): Boolean {
        var sum = 0
        for (i in 1..num/2) {
            if (num % i == 0) {
                sum += i
            }
        }
        return num == sum
    }
}
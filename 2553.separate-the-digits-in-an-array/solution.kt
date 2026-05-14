class Solution {
    fun separateDigits(nums: IntArray): IntArray {
        val answer = mutableListOf<Int>()

        fun addDigits(num: Int) {
            if (num > 0) {
                addDigits(num / 10)
                answer += num % 10
            }
        }

        for (num in nums) {
            addDigits(num)
        }
        
        return answer.toIntArray()
    }
}
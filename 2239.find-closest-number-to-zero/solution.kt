import kotlin.math.abs
import kotlin.math.sign

class Solution {
    fun findClosestNumber(nums: IntArray): Int {
        return nums.reduce { acc, num ->
            when ((abs(num) - abs(acc)).sign) {
                -1 -> num
                0 -> maxOf(acc, num)
                else -> acc
            }
        }
    }
}
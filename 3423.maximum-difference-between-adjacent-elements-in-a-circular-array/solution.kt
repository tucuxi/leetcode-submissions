class Solution {
    fun maxAdjacentDistance(nums: IntArray): Int =
        nums.fold(nums.last() to 0) { (prev, diff), num ->
            num to maxOf(diff, num - prev, prev - num)
        }.second
}
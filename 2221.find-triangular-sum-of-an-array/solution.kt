class Solution {
    fun triangularSum(nums: IntArray): Int = rec(nums.asList())

    private tailrec fun rec(nums: List<Int>): Int =
        if (nums.size == 1) {
            nums.first()
        } else {
            rec(
                nums
                    .zipWithNext()
                    .map { (a, b) -> (a + b) % 10 }
                )
        }
}
class Solution {
    fun minimumPairRemoval(nums: IntArray): Int {
        val list = LinkedList(nums.asList())

        while (list.size > 1 && !nonDecreasing(list)) {
            var minSum = Int.MAX_VALUE
            var minSumIndex = 0
            
            (0 until list.lastIndex).forEach { i ->
                val sum = list[i] + list[i + 1]
                if (sum < minSum) {
                    minSum = sum
                    minSumIndex = i
                }
            }

            list.removeAt(minSumIndex)
            list.set(minSumIndex, minSum)
        }

        return nums.size - list.size
    }

    fun <T : Comparable<T>> nonDecreasing(list: List<out T>): Boolean =
        (0 until list.lastIndex).all { i -> list[i] <= list[i + 1] }
}
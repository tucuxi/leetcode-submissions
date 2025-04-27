class Solution {
    fun smallestRange(nums: List<List<Int>>): IntArray {
        val allNumsWithIndex = nums
            .flatMapIndexed { index, list -> list.map { num -> num to index } }
            .sortedBy { (num, _) -> num }
        val currentNumPerList = Array(nums.size) { Int.MIN_VALUE }
        val minHeap = PriorityQueue<Int>()
        val maxHeap = PriorityQueue<Int> { a, b -> b - a }
        var numListsVisited = 0
        var a = -100001
        var b = 100001
        for ((num, index) in allNumsWithIndex) {
            val currentNum = currentNumPerList[index]
            if (currentNum == Int.MIN_VALUE) {
                numListsVisited++
            } else {
                minHeap.remove(currentNum)
                maxHeap.remove(currentNum)
            }
            currentNumPerList[index] = num
            minHeap.add(num)
            maxHeap.add(num)
            if (numListsVisited == nums.size) {
                val c = minHeap.peek()
                val d = maxHeap.peek()
                if (d - c < b - a) {
                    a = c
                    b = d
                }
            }
        }
        return intArrayOf(a, b)
    }
}
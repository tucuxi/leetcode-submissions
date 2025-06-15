class Solution {
    fun furthestBuilding(heights: IntArray, bricks: Int, ladders: Int): Int {
        var remainingBricks = bricks
        val pq = PriorityQueue<Int>()
        val leftOver = heights.asSequence()
            .zipWithNext()
            .dropWhile { (firstHeight, secondHeight) ->
                if (firstHeight < secondHeight) {
                    pq.add(secondHeight - firstHeight)
                    if (pq.size > ladders) {
                        remainingBricks -= pq.remove()
                    }
                }
                
                remainingBricks >= 0
            }
            .toList()
        return heights.lastIndex - leftOver.size
    }
}
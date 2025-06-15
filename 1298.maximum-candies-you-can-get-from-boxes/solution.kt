class Solution {
    fun maxCandies(status: IntArray, candies: IntArray, keys: Array<IntArray>, containedBoxes: Array<IntArray>, initialBoxes: IntArray): Int {
        var res = 0
        val boxes = BooleanArray(status.size)
        val queue = ArrayDeque<Int>()

        initialBoxes.forEach { i ->
            if (status[i] == 0) {
                boxes[i] = true
            } else {
                queue.add(i)
            }
        }
        
        while (queue.isNotEmpty()) {
            val i = queue.removeFirst()
            res += candies[i]

            containedBoxes[i].forEach { j ->
                if (status[j] == 0) {
                    boxes[j] = true
                } else {
                    queue.add(j)
                }
            }

            keys[i].forEach { j ->
                if (status[j] == 0) {
                    if (boxes[j]) {
                        boxes[j] = false
                        queue.add(j)
                    } else {
                        status[j] = 1
                    }
                }
            }
        }
        return res
    }
}
import java.util.SortedMap

class Solution {
    fun minimumCost(nums: IntArray, k: Int, dist: Int): Long {
        val needed = k - 1
        var currentSum = 0L
        var result = Long.MAX_VALUE
        
        val top = sortedMapOf<Int, Int>()
        val rest = sortedMapOf<Int, Int>()
        var topCount = 0

        fun SortedMap<Int, Int>.addOne(num: Int) {
            this[num] = this.getOrDefault(num, 0) + 1
        }

        fun SortedMap<Int, Int>.removeOne(num: Int) {
            val count = this[num] ?: return
            if (count == 1) this.remove(num) else this[num] = count - 1
        }

        for (i in 1 until nums.size) {
            val num = nums[i]
            rest.addOne(num)

            if (topCount < needed || (top.isNotEmpty() && num < top.lastKey())) {
                val smallestRest = rest.firstKey()
                rest.removeOne(smallestRest)
                top.addOne(smallestRest)                
                currentSum += smallestRest
                topCount++
            }

            if (topCount > needed) {
                val largestTop = top.lastKey()
                top.removeOne(largestTop)
                rest.addOne(largestTop)
                currentSum -= largestTop
                topCount--
            }

            if (i > dist + 1) {
                val outgoing = nums[i - dist - 1]
                if (top.containsKey(outgoing)) {
                    top.removeOne(outgoing)
                    currentSum -= outgoing
                    topCount--
                } else {
                    rest.removeOne(outgoing)
                }
            }

            if (topCount < needed && rest.isNotEmpty()) {
                val smallestRest = rest.firstKey()
                rest.removeOne(smallestRest)
                top.addOne(smallestRest)
                currentSum += smallestRest
                topCount++
            }

            if (i >= needed) {
                result = minOf(result, currentSum)
            }
        }

        return nums[0] + result
    }
}
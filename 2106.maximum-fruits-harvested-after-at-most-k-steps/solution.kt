class Solution {
    fun maxTotalFruits(fruits: Array<IntArray>, startPos: Int, k: Int): Int {
        
        fun steps(left: Int, right: Int): Int =
            when {
                fruits[right][0] <= startPos ->
                    startPos - fruits[left][0]
                
                fruits[left][0] >= startPos ->
                    fruits[right][0] - startPos
            
                else ->
                    fruits[right][0] -
                    fruits[left][0] +
                    minOf(
                        abs(startPos - fruits[right][0]),
                        abs(startPos - fruits[left][0]) 
                    )
            }

        var left = fruits
            .asList()
            .binarySearchBy(startPos - k) { it[0] }
            .let { if (it < 0) -(it + 1) else it }
    
        var right = left
        var sum = 0
        var res = 0

        while (right < fruits.size && fruits[right][0] <= startPos + k) {
            if (steps(left, right) > k) {
                sum -= fruits[left][1]
                left++
            } else {
                sum += fruits[right][1]
                res = maxOf(sum, res)
                right++
            }
        }
        
        return res
    }

}

/*
Case 1: right is left of startPos or equal

startPos - leftPos

Case 2: right is right of startPos

rightPos - startPos

Case 3: left is left of startPos, right is right of startPos

(startPos - left) + (right - left) = startPos + right - 2 * left
*/
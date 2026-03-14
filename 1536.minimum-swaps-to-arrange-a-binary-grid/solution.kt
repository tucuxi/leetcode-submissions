class Solution {
    fun minSwaps(grid: Array<IntArray>): Int {
        val n = grid.size

        val a = IntArray(grid.size) { i ->
            var j = grid.lastIndex
            
            while (j >= 0 && grid[i][j] == 0) {
                j--
            }
            
            n - 1 - j
        }

        var swaps = 0
        
        for (i in a.indices) {
            var k = -1
            for (j in i until a.size) {
                if (a[j] >= n - 1 - i) {
                    swaps += j - i
                    k = j
                    break
                }
            }
            if (k == -1) {
                return -1
            }
            for (j in k downTo i + 1) {
                a[j] = a[j - 1].also { a[j - 1] = a[j] }
            }
        }

        return swaps
    }
}
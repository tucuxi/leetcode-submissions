class Solution {
    fun numOfUnplacedFruits(fruits: IntArray, baskets: IntArray): Int {
        var res = 0
        val v = BooleanArray(baskets.size)

        fruits.forEach(
            fun(f: Int) {
                baskets.forEachIndexed { i, b ->
                    if (!v[i] && b >= f) {
                        v[i] = true
                        return
                    }            
                }
                res++
            }
        )

        return res
    }
}
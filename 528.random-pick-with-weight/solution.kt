class Solution(w: IntArray) {

    private val prefixSums = mutableListOf<Int>()
    private var totalSum = 0

    init {
        for (weight in w) {
            totalSum += weight
            prefixSums.add(totalSum)
        }
    }

    fun pickIndex(): Int {
        val target = (1..totalSum).random()

        // binarySearch returns the index of the element if found.
        // If not found, it returns (-(insertion point) - 1).
        val index = prefixSums.binarySearch(target)
        
        return if (index >= 0) index else -(index + 1)
    }

}

/**
 * Your Solution object will be instantiated and called as such:
 * var obj = Solution(w)
 * var param_1 = obj.pickIndex()
 */
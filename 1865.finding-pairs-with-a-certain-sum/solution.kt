class FindSumPairs(nums1: IntArray, nums2: IntArray) {

    private val h1 = nums1.asSequence().groupingBy { it }.eachCount()
    private val h2 = nums2.copyOf()

    fun add(index: Int, `val`: Int) {
        h2[index] += `val`    
    }

    fun count(tot: Int): Int {
        return h2.fold(0) { acc, num2 -> acc + h1.getOrDefault(tot - num2, 0) }
    }

}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * var obj = FindSumPairs(nums1, nums2)
 * obj.add(index,`val`)
 * var param_2 = obj.count(tot)
 */
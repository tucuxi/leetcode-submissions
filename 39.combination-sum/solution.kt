class Solution {
    fun combinationSum(candidates: IntArray, target: Int): List<List<Int>> {
        val result = mutableSetOf<List<Int>>()
        val combination = mutableListOf<Int>()
        
        fun recurse(remainder: Int) {
            if (remainder == 0) { 
                result.add(combination.sorted())
                return
            }
            for (n in candidates) {
                if (n <= remainder) {
                    combination.add(n)
                    recurse(remainder - n)
                    combination.removeAt(combination.lastIndex)
                }
            }
        }
        
        recurse(target)
        return result.toList()
    }
}
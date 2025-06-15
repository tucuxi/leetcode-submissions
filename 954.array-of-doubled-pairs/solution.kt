class Solution {
    fun canReorderDoubled(arr: IntArray): Boolean {
        val frequencies = mutableMapOf<Int, Int>()
        arr.forEach {
            frequencies[it] = (frequencies[it] ?: 0) + 1
        }        
        
        for (num in arr.sorted()) {
            if (frequencies[num]!! > 0) {
                val doubleNum = 2 * num
                val doubleNumFrequency = frequencies.getOrDefault(doubleNum, 0)
                if (doubleNumFrequency > 0) {
                    frequencies[doubleNum] = doubleNumFrequency - 1
                    frequencies[num] = frequencies[num]!! - 1
                }
            }
        }
        return frequencies.values.none { it != 0 }
    }
}
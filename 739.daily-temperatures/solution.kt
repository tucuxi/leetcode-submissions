class Solution {
    fun dailyTemperatures(temperatures: IntArray): IntArray {
        var res = IntArray(temperatures.size) { 0 }
        val s = ArrayDeque<Int>(temperatures.size)
        for (index in temperatures.indices) {
            while (s.isNotEmpty() && temperatures[index] > temperatures[s.last()]) {
                val day = s.removeLast()
                res[day] = index - day
            }
            s.add(index)
        }
        return res
    }
}
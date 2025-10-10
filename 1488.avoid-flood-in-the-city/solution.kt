class Solution {
    fun avoidFlood(rains: IntArray): IntArray {
        val ans = IntArray(rains.size) { -1 }
        val previousRainyDayForLake = mutableMapOf<Int, Int>()
        val dryDays = TreeSet<Int>()

        for ((day, lake) in rains.withIndex()) {
            if (lake == 0) {
                dryDays.add(day)
                ans[day] = 1
            } else {
                previousRainyDayForLake[lake]?.let { 
                    val drainDay = dryDays.higher(it) ?: return intArrayOf()
                    dryDays.remove(drainDay)
                    ans[drainDay] = lake
                }
                previousRainyDayForLake[lake] = day
            }
        }

        return ans
    }
}
class Solution {
    fun mergeTriplets(triplets: Array<IntArray>, target: IntArray): Boolean {
        val reachable = BooleanArray(3)
        triplets
            .filter { t -> allLessOrEqual(t, target) }
            .forEach { t ->
                reachable
                    .forEachIndexed { i, r ->
                        reachable[i] = r || t[i] == target[i]
                    }
            }
        return reachable.all { it }
    }

    fun allLessOrEqual(a: IntArray, b: IntArray): Boolean =
        (a zip b).all { it.first <= it.second }
}
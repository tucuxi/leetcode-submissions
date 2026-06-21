class Solution {
    fun asteroidsDestroyed(mass: Int, asteroids: IntArray): Boolean {
        var m = mass.toLong()
        asteroids.sort()
        for (a in asteroids) {
            if (m < a) {
                return false
            }
            m += a
        }
        return true
    }
}
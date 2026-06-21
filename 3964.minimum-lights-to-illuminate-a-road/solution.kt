class Solution {
    fun minLights(lights: IntArray): Int {
        val n = lights.size
        val vis = BooleanArray(n)

        lights.forEachIndexed { i, v ->
            if (v > 0) {
                for (j in max(0, i - v) .. min(n - 1, i + v)) {
                    vis[j] = true
                }
            }
        }

        var res = 0
        var i = 0

        while (i < n) {
            if (!vis[i]) {
                res++
                i += 2
            }
            i++
        }

        return res
    }
}
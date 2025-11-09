class Solution {
    fun maxTaskAssign(tasks: IntArray, workers: IntArray, pills: Int, strength: Int): Int {
        tasks.sort()
        workers.sort()

        var low = 0
        var high = tasks.lastIndex
        var res = -1

        while (low <= high) {
            val m = low + (high - low) / 2
            var j = workers.lastIndex
            var p = pills
            var good = true
            val q = mutableListOf<Int>()
            for (i in m downTo 0) {
                while (j >= maxOf(0, workers.lastIndex - m) && workers[j] + strength >= tasks[i]) {
                    q += workers[j]
                    j--
                }
                if (q.isEmpty()) {
                    good = false
                    break
                }
                if (q.first() >= tasks[i]) {
                    q.removeFirst()
                } else if (--p < 0) {
                    good = false
                    break
                } else {
                    q.removeLast()
                }
            }
            if (good) {
                res = maxOf(m, res)
                low = m + 1
            } else {
                high = m - 1
            }
        }
        return res + 1
    }
}
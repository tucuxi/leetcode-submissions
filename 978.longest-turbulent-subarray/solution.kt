class Solution {
    fun maxTurbulenceSize(arr: IntArray): Int {
        var s = 0
        var run = 1
        var max = 1
        for (i in 1..arr.lastIndex) {
            if (arr[i] < arr[i-1]) {
                if (s < 0) {
                    run = 1
                }
                s = -1
            } else if (arr[i] > arr[i-1]) {
                if (s > 0) {
                    run = 1
                }
                s = 1
            } else {
                run = 0
                s = 0
            }
            max = maxOf(max, ++run)
        }
        return max
    }
}
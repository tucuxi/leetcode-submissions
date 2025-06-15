class Solution {
    fun distanceBetweenBusStops(distance: IntArray, start: Int, destination: Int): Int {
        var sum = 0
        var i = start
        while (i != destination) {
            sum += distance[i]
            if (++i == distance.size)
                i = 0
        }
        return Math.min(sum, distance.sum() - sum)
    }
}
class UndergroundSystem() {

    private val trip = mutableMapOf<Int, Pair<String, Int>>()
    private val total = mutableMapOf<Pair<String, String>, Pair<Int, Long>>()
    
    fun checkIn(id: Int, stationName: String, t: Int) {
        trip[id] = stationName to t
    }

    fun checkOut(id: Int, stationName: String, t: Int) {
        trip[id]?.let { (start, t1) ->
            val (n, tSum) = total[start to stationName] ?: 0 to 0L
            total[start to stationName] = (n + 1) to (tSum + t - t1)
        }
    }

    fun getAverageTime(startStation: String, endStation: String): Double =
        total[startStation to endStation]?.let {
            (n, tSum) -> tSum.toDouble() / n
        } ?: 0.0

}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * var obj = UndergroundSystem()
 * obj.checkIn(id,stationName,t)
 * obj.checkOut(id,stationName,t)
 * var param_3 = obj.getAverageTime(startStation,endStation)
 */
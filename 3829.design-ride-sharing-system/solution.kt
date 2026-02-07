class RideSharingSystem() {

    private val riders = mutableListOf<Int>()
    private val drivers = mutableListOf<Int>()
    private val waiting = mutableSetOf<Int>()

    fun addRider(riderId: Int) {
        riders += riderId
        waiting += riderId
    }

    fun addDriver(driverId: Int) {
        drivers += driverId
    }

    fun matchDriverWithRider(): IntArray {
        removeCanceledRides()
        return if (drivers.isEmpty() || riders.isEmpty()) {
            intArrayOf(-1, -1)
        } else {
            intArrayOf(drivers.removeFirst(), riders.removeFirst())
        }
    }

    fun cancelRider(riderId: Int) {
        waiting -= riderId
    }
    
    private fun removeCanceledRides() {
        while (riders.isNotEmpty() && riders.first() !in waiting) {
            waiting -= riders.removeFirst()
        }
    }
}

/**
 * Your RideSharingSystem object will be instantiated and called as such:
 * var obj = RideSharingSystem()
 * obj.addRider(riderId)
 * obj.addDriver(driverId)
 * var param_3 = obj.matchDriverWithRider()
 * obj.cancelRider(riderId)
 */
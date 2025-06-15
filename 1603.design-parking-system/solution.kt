class ParkingSystem(big: Int, medium: Int, small: Int) {

    private var freeBig = big
    private var freeMedium = medium
    private var freeSmall = small
    
    fun addCar(carType: Int): Boolean {
        val res = when (carType) {
            1 -> freeBig > 0
            2 -> freeMedium > 0
            3 -> freeSmall > 0
            else -> false
        }
        if (res) when (carType) {
            1 -> freeBig -= 1
            2 -> freeMedium -= 1
            3 -> freeSmall -= 1
        }
        return res
    }

}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * var obj = ParkingSystem(big, medium, small)
 * var param_1 = obj.addCar(carType)
 */
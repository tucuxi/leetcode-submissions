type ParkingSystem struct {
    free []int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{[]int{0, big, medium, small}}
}


func (this *ParkingSystem) AddCar(carType int) bool {
    res := this.free[carType] > 0
    if (res) {
        this.free[carType]--
    }
    return res
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
type ParkingSystem struct {
	curCount [3]int
}
func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		curCount: [3]int{big, medium, small},
	}
}
func (this *ParkingSystem) AddCar(carType int) bool {
	if this.curCount[carType-1] == 0 {
		return false
	}
	this.curCount[carType-1]--
	return true
}

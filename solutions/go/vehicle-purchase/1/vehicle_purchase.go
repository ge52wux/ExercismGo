package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    var isCar bool = (kind == "car") 
    var isTruck bool = (kind == "truck")
	return isCar || isTruck
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    var result string 
	if option1 <= option2 {
        result =  option1
    } else {
        result = option2
    }
    return result + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    var priceCoefficient float64 = 1
    if age >= 10 {
        priceCoefficient = 0.5
    } else if age >= 3 {
        priceCoefficient = 0.7
    } else {
        priceCoefficient = 0.8
    }	
    return float64(originalPrice)*priceCoefficient
}

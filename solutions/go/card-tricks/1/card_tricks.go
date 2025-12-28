package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	var result int 
    if len(slice) <= index  || index < 0 {
        result = -1
    } else {
        result = slice[index]
    }
    return result
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    result := slice
	if len(slice) <= index  || index < 0 {
        result = append(result, value)
    } else {
        result[index] = value
    }
    return result
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    var result []int
    if len(slice) <= index  || index < 0 {
        result = slice
    } else {
		result = append(slice[:index], slice[index+1:]...)
    }
    return result
}

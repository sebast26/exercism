package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
	units := make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill create a new bill
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem add item to customer bill
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, ok := units[unit]
	if !ok {
		return false
	}
	bill[item] = quantity
	return true
}

// RemoveItem remove item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	quantity, billOk := bill[item]
	u, unitsOk := units[unit]
	if !billOk || !unitsOk {
		return false
	}
	newQuantity := quantity - u
	if newQuantity < 0 {
		return false
	}
	if newQuantity == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = newQuantity
	return true
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, ok := bill[item]
	if !ok {
		return 0, false
	}
	return quantity, true
}

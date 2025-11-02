package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	number, exists := units[unit]
    if exists {
        bill[item] += number
        return true
    }
    return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	number, validUnit := units[unit]
    quantity, inBill := bill[item]
    switch {
    case !validUnit:
        return false
	case !inBill:
        return false    
    case number > quantity:
        return false
    case number == quantity:
    	delete(bill, item)
        return true
    case quantity > number:
    	bill[item] = bill[item] - number
        return true
    default:
    	return false
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	number, exists := bill[item]
    if exists {
        return number, true
    }
	return 0, false
}
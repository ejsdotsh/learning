package lasagna

// PreparationTime accepts a list of ingredients and an average prep-time per-layer,
// and returns how log it will take to prepare.
func PreparationTime(ingredients []string, avgPrepTime int) int {
    if avgPrepTime < 1 {
        return 2 * len(ingredients)
    }
	return avgPrepTime * len(ingredients)
}

// Quantities takes a slice of layers ([]string) and returns the amount of noodles (int) and sauce (float64) required.
func Quantities(ingredients []string) (noodles int, sauce float64) {
    for _, v := range ingredients {
        if v == "noodles" {
            noodles += int(50)
        }
    	if v == "sauce" {
            sauce += float64(0.2)
        }
    }
	return
}

// AddSecretIngredient updates myList with the last entry in friendsList and returns nothing.
func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// ScaleRecipt returns the quantities needed to scale the two portion recipe.
func ScaleRecipe(quantities []float64, portions int) []float64 {
    var mp float64 = float64(portions) / float64(2)
    uq := make([]float64, 0)
    for _, v := range quantities {
         uq = append(uq, mp * v)
    }
	return uq
}
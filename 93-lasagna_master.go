package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTPerLayer int) int {
    if avgTPerLayer == 0 {
        avgTPerLayer = 2
    }
    return len(layers) * avgTPerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodleQ := 0
    sauceQ := 0.00
    
	for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodleQ += 50
        } else if layers[i] == "sauce" {
            sauceQ += 0.2
        }
    }
    return noodleQ, sauceQ
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
    myList[len(myList) - 1] = friendList[len(friendList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64{
    scaleQ := make([]float64, len(quantities))
    for i := 0; i < len(quantities); i++ {
        scaleQ[i] = quantities[i] * (float64(portions) / 2)
    }
	return scaleQ
}

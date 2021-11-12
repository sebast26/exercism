package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0
	for _, l := range layers {
		if l == "noodles" {
			noodles++
		}
		if l == "sauce" {
			sauce++
		}
	}
	return noodles * 50, 0.2 * float64(sauce)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	s := make([]float64, len(quantities))
	for i := range quantities {
		s[i] = quantities[i] * float64(portions) / 2
	}
	return s
}

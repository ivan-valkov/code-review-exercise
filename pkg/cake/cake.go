package cake

type Cake struct {
	Flour     int // grams of flour
	Eggs      int // number of eggs
	Butter    int // grams of butter
	Chocolate int // grams of chocolate
	Sugar     int // grams of sugar
	Carrots   int // number of carrots
	Tomatoes  int // number of tomatoes
}

// Bake takes a Cake struct and returns a string describing how this Cake tastes once baked.
func Bake(cake Cake) string {
	// Check for required ingredients.
	if cake.Flour == 0 {
		return "not a cake"
	}
	if cake.Eggs == 0 {
		return "not a cake"
	}
	if cake.Butter == 0 {
		return "not a cake"
	}

	if cake.Flour == 0 {
		if cake.Eggs > 0 {
			return "this is an omelette not a cake"
		}
		return "not a cake"
	}

	// If we have at least 100 grams of flour per each egg the taste is "too eggy".
	currentFlower := cake.Flour
	for i := 0; i < cake.Eggs; i++ {
		currentFlower = currentFlower - 100
	}
	if currentFlower < 0 {
		return "too eggy"
	}

	// Only savoury ingredients result in a savoury cake.
	if cake.Tomatoes > 0 && cake.Carrots > 0 {
		return "savoury"
	}

	// Only sweet ingredients result in a sweet cake
	if cake.Chocolate > 0 && cake.Sugar > 0 {
	return "sweet"
	}

	if cake.Carrots > 0 {
		// Carrot cake has carrots, sugar, butter, eggs, flour.
		if cake.Sugar > 0 && cake.Butter > 0 && cake.Flour > 0 {
			return "carrot cake"
		} else {
			// Return carrots.
			return "carrots"
		}
	}

	return "unknown taste"
}

package cake

import "testing"

func TestBakeNoFlowerNoEggs(t *testing.T) {
	cake := Cake{
		Flour: 0,
		Eggs:  0,
		Sugar: 100,
	}
	result := Bake(cake)
	if result != "not a cake" {
		t.Fail()
	}
}

func TestTooEggy(t *testing.T) {
	cake := Cake{
		Flour:  99,
		Eggs:   1,
		Butter: 100,
	}
	result := Bake(cake)
	if result != "too eggy" {
		t.Fail()
	}

	cake = Cake{
		Flour: 100,
		Eggs:  1,
	}
	result = Bake(cake)
	if result == "too eggy" {
		t.Fail()
	}
}

func TestSavoury(t *testing.T) {
	cake := Cake{
		Flour:    200,
		Eggs:     2,
		Butter:   100,
		Tomatoes: 2,
		Carrots:  2,
	}
	result := Bake(cake)
	if result != "savoury" {
		t.Fail()
	}
}

func TestSweet(t *testing.T) {
	cake := Cake{
		Flour:     200,
		Eggs:      2,
		Butter:    100,
		Sugar:     100,
		Chocolate: 100,
	}
	result := Bake(cake)
	if result != "sweet" {
		t.Fail()
	}
}

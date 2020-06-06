package dice

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestDicePoolSum(t *testing.T) {
	testCases := map[int]dicePool{
		12: dicePool{rolled: []int{3, 4, 5}},
		20: dicePool{rolled: []int{12, 1, 7}},
		22: dicePool{rolled: []int{12, 1, 7}, modifier: 2},
		18: dicePool{rolled: []int{12, 1, 7}, modifier: -2},
	}

	for expected, dp := range testCases {
		if expected != dp.Sum() {
			t.Logf("TestDicePoolSum: expected %d, got %d",
				expected, dp.Sum())
			t.Fail()
		}
	}
}

func TestDicePoolNormalRoll(t *testing.T) {
	dp := dicePool{numDice: 3, numSides: 6}
	exp := []int{3, 5, 6} // Sorted rolls
	dp.Roll(getRNG())
	if reflect.DeepEqual(dp.rolled, exp) == false {
		t.Logf("Normal roll failed, expected %v, got %v", exp, dp.rolled)
		t.Fail()
	}
}

func TestDicePoolExplodingRoll(t *testing.T) {

	dp := dicePool{numDice: 4, numSides: 6, exploding: true}
	exp := []int{3, 3, 5, 5, 6, 6} // Sorted rolls
	dp.Roll(getRNG())
	if reflect.DeepEqual(dp.rolled, exp) == false {
		t.Logf("Exploding roll failed, expected %v, got %v", exp, dp.rolled)
		t.Fail()
	}
}

func getRNG() *rand.Rand {
	// For the given seed of 1337, the following numbers will be
	// generated when rolling d6 dice
	//d6 := []int{4, 2, 5, 5, 2, 4, 4, 5, 1, 5, 0, 0, 5, 3, 2}

	source := rand.NewSource(1337)
	return rand.New(source)
}

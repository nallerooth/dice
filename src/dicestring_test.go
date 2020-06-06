package dice

import (
	"testing"
)

func TestParseDiceString(t *testing.T) {
	var err error
	var dp *dicePool

	testCases := map[string]dicePool{
		"3d6":    dicePool{numDice: 3, numSides: 6},
		"12d6":   dicePool{numDice: 12, numSides: 6},
		"4d2":    dicePool{numDice: 4, numSides: 2},
		"5d20":   dicePool{numDice: 5, numSides: 20},
		"3d6+0":  dicePool{numDice: 3, numSides: 6},
		"3d6-0":  dicePool{numDice: 3, numSides: 6},
		"3d6+2":  dicePool{numDice: 3, numSides: 6, modifier: 2},
		"3d6-2":  dicePool{numDice: 3, numSides: 6, modifier: -2},
		"3d6!":   dicePool{numDice: 3, numSides: 6, exploding: true},
		"3d6+1!": dicePool{numDice: 3, numSides: 6, modifier: 1, exploding: true},
	}

	for dicestr, expected := range testCases {
		if dp, err = parseDiceString(dicestr); err != nil {
			t.Log(err)
			t.Fail()
		}

		if dp.numDice != expected.numDice {
			t.Logf("Test case [%s] failed: Expected %d dice, got %d\n",
				dicestr, expected.numDice, dp.numDice)
			t.Fail()
		}

		if dp.numSides != expected.numSides {
			t.Logf("Test case [%s] failed: Expected %d sides, got %d\n",
				dicestr, expected.numSides, dp.numSides)
			t.Fail()
		}

		if dp.modifier != expected.modifier {
			t.Logf("Test case [%s] failed: Expected modifier %+d, got %+d\n",
				dicestr, expected.modifier, dp.modifier)
			t.Fail()
		}

		if dp.exploding != expected.exploding {
			t.Logf("Test case [%s] failed: Expected exploding %t, got %t\n",
				dicestr, expected.exploding, dp.exploding)
			t.Fail()
		}
	}

}

func TestMalformedDiceString(t *testing.T) {
	testCases := []string{
		"-387f738",
		"dfsd",
		"3dd6",
		"3d-4",
		"3d6_3",
		"3d6+d",
	}

	for _, brokenstring := range testCases {
		if _, err := parseDiceString(brokenstring); err == nil {
			t.Logf("Broken dicestring <%s> should throw an error", brokenstring)
			t.Fail()
		}
	}
}

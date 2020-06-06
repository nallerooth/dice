package dice

import (
	"testing"
)

func TestParseDiceString(t *testing.T) {
	var err error
	var dp *DicePool

	testCases := map[string]DicePool{
		"3d6":    DicePool{NumDice: 3, NumSides: 6},
		"12d6":   DicePool{NumDice: 12, NumSides: 6},
		"4d2":    DicePool{NumDice: 4, NumSides: 2},
		"5d20":   DicePool{NumDice: 5, NumSides: 20},
		"3d6+0":  DicePool{NumDice: 3, NumSides: 6},
		"3d6-0":  DicePool{NumDice: 3, NumSides: 6},
		"3d6+2":  DicePool{NumDice: 3, NumSides: 6, Modifier: 2},
		"3d6-2":  DicePool{NumDice: 3, NumSides: 6, Modifier: -2},
		"3d6!":   DicePool{NumDice: 3, NumSides: 6, Exploding: true},
		"3d6+1!": DicePool{NumDice: 3, NumSides: 6, Modifier: 1, Exploding: true},
	}

	for dicestr, expected := range testCases {
		if dp, err = parseDiceString(dicestr); err != nil {
			t.Log(err)
			t.Fail()
		}

		if dp.NumDice != expected.NumDice {
			t.Logf("Test case [%s] failed: Expected %d dice, got %d\n",
				dicestr, expected.NumDice, dp.NumDice)
			t.Fail()
		}

		if dp.NumSides != expected.NumSides {
			t.Logf("Test case [%s] failed: Expected %d sides, got %d\n",
				dicestr, expected.NumSides, dp.NumSides)
			t.Fail()
		}

		if dp.Modifier != expected.Modifier {
			t.Logf("Test case [%s] failed: Expected modifier %+d, got %+d\n",
				dicestr, expected.Modifier, dp.Modifier)
			t.Fail()
		}

		if dp.Exploding != expected.Exploding {
			t.Logf("Test case [%s] failed: Expected exploding %t, got %t\n",
				dicestr, expected.Exploding, dp.Exploding)
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

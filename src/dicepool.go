package dice

import (
	"math/rand"
	"sort"
)

// roll holds a number of a given die along with any modifier to apply
type DicePool struct {
	Rolled    []int
	NumSides  int
	NumDice   int
	Modifier  int
	Exploding bool
}

// Sum will add the values from any rolled dice and then add the modifier
// Note: a negative modifier can make the result negative
func (r *DicePool) Sum() int {
	var sum int

	for _, d := range r.Rolled {
		sum += d
	}

	return sum + r.Modifier
}

// roll will append the result from each die to the dicePool::rolled list
func (dp *DicePool) roll(rng *rand.Rand) {
	for i := 0; i < dp.NumDice; i++ {

		if dp.Exploding {
			for {
				r := rng.Intn(dp.NumSides) + 1
				dp.Rolled = append(dp.Rolled, r)

				if r != dp.NumSides {
					break
				}
			}
		} else {
			dp.Rolled = append(dp.Rolled, rng.Intn(dp.NumSides)+1)
		}
	}

	sort.Ints(dp.Rolled)
}

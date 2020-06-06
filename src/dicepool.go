package dice

import (
	"math/rand"
	"sort"
)

// roll holds a number of a given die along with any modifier to apply
type dicePool struct {
	rolled    []int
	numSides  int
	numDice   int
	modifier  int
	exploding bool
}

// Sum will add the values from any rolled dice and then add the modifier
// Note: a negative modifier can make the result negative
func (r *dicePool) Sum() int {
	var sum int

	for _, d := range r.rolled {
		sum += d
	}

	return sum + r.modifier
}

// Roll will append the result from each die to the dicePool::rolled list
func (dp *dicePool) Roll(rng *rand.Rand) {
	for i := 0; i < dp.numDice; i++ {

		if dp.exploding {
			for {
				r := rng.Intn(dp.numSides) + 1
				dp.rolled = append(dp.rolled, r)

				if r != dp.numSides {
					break
				}
			}
		} else {
			dp.rolled = append(dp.rolled, rng.Intn(dp.numSides)+1)
		}
	}

	sort.Ints(dp.rolled)
}

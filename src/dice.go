package dice

import (
	"fmt"
	"math/rand"
	"time"
)

// Dice holds a RNG for future rolls
type Dice struct {
	rng *rand.Rand
}

// New creates a new Dice object with its own RNG
func New() *Dice {
	d := Dice{}
	source := rand.NewSource(time.Now().UnixNano())
	d.rng = rand.New(source)

	return &d
}

// Roll will take a DiceString, parse it and then proceed to make the
// proper rolls, returning a result
func (d *Dice) RollDetailed(dicestr string) (*dicePool, error) {
	var dp *dicePool
	var err error

	if dp, err = parseDiceString(dicestr); err != nil {
		return nil, fmt.Errorf("Dice Error: %s\n", err)
	}

	dp.Roll(d.rng)

	return dp, nil
}

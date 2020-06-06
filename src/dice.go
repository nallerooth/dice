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

// RollDetailed takes a DiceString and returns a dicePool
func (d *Dice) RollDetailed(dicestr string) (*DicePool, error) {
	var dp *DicePool
	var err error

	if dp, err = parseDiceString(dicestr); err != nil {
		return nil, fmt.Errorf("Dice Error: %s\n", err)
	}

	dp.roll(d.rng)

	return dp, nil
}

// RollSum takes a DiceString and returns the sum of the dice,
// including any modifiers
func (d *Dice) RollSum(dicestr string) (int, error) {
	dp, err := d.RollDetailed(dicestr)
	if err != nil {
		return 0, err
	}

	return dp.Sum(), nil
}

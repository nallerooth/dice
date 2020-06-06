package dice

import (
	"fmt"
	"regexp"
	"strconv"
)

// parseDiceString will attempt to break down a string into useful bits
// Examples:
//   "3d6" -> 3 x D6
//   "3d6+2" -> 3 x D6 and then append 2 to the sum of all the dice
func parseDiceString(dicestr string) (*DicePool, error) {
	var err error

	pattern := regexp.MustCompile(`^(\d+)d(\d+)([+-]?)(\d*)(!?)$`)
	matches := pattern.FindStringSubmatch(dicestr)

	dp := DicePool{}

	if len(matches) < 2 {
		return nil, fmt.Errorf("Malformed sice string: %s\n", dicestr)
	}

	if dp.NumDice, err = strconv.Atoi(matches[1]); err != nil {
		return nil, err
	}

	if dp.NumSides, err = strconv.Atoi(matches[2]); err != nil {
		return nil, err
	}

	if matches[4] != "" {
		if dp.Modifier, err = strconv.Atoi(matches[4]); err != nil {
			return nil, err
		}
		if matches[3] == "-" {
			dp.Modifier *= -1
		}
	}

	if matches[5] == "!" {
		dp.Exploding = true
	}

	return &dp, nil
}

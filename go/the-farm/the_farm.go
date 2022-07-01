package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	fodder, err := weightFodder.FodderAmount()
	if err != nil {
		if err == ErrScaleMalfunction && fodder > 0 {
			return fodder * 2 / float64(cows), nil
		}
		if err == ErrScaleMalfunction && fodder < 0 {
			return 0, errors.New("negative fodder")
		}
		return 0, err
	}
	if fodder < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	}
	return fodder / float64(cows), err
}

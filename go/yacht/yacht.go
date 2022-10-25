package yacht

import "sort"

type scoring func([]int) int

var scores = map[string]scoring{
	"ones":            scoring(func(dice []int) int { return sumIf(dice, 1) }),
	"twos":            scoring(func(dice []int) int { return sumIf(dice, 2) }),
	"threes":          scoring(func(dice []int) int { return sumIf(dice, 3) }),
	"fours":           scoring(func(dice []int) int { return sumIf(dice, 4) }),
	"fives":           scoring(func(dice []int) int { return sumIf(dice, 5) }),
	"sixes":           scoring(func(dice []int) int { return sumIf(dice, 6) }),
	"full house":      scoreFullHouse,
	"four of a kind":  scoreFourOfKind,
	"little straight": scoreLittleStraight,
	"big straight":    scoreBigStraight,
	"choice":          scoring(func(dice []int) int { return sum(dice) }),
	"yacht":           scoreYacht,
}

func Score(dice []int, category string) int {
	s, ok := scores[category]
	if !ok {
		return 0
	}
	return s(dice)
}

var scoreFullHouse = scoring(func(dice []int) int {
	sort.Ints(dice)

	if !((allEq(dice[0:2]) && allEq(dice[2:5])) || (allEq(dice[0:3]) && allEq(dice[3:5]))) || allEq(dice) {
		return 0
	}

	return sum(dice)
})

var scoreFourOfKind = scoring(func(dice []int) int {
	sort.Ints(dice)

	if !((allEq(dice[0:4]) && allEq(dice[4:5])) || (allEq(dice[0:1]) && allEq(dice[2:5]))) {
		return 0
	}

	if allEq(dice[0:4]) {
		return sum(dice[0:4])
	} else {
		return sum(dice[1:5])
	}
})

var scoreLittleStraight = scoring(func(dice []int) int {
	sort.Ints(dice)

	if allIncFrom(dice, 1) {
		return 30
	} else {
		return 0
	}
})

var scoreBigStraight = scoring(func(dice []int) int {
	sort.Ints(dice)

	if allIncFrom(dice, 2) {
		return 30
	} else {
		return 0
	}
})

var scoreYacht = scoring(func(dice []int) int {
	if allEq(dice) {
		return 50
	} else {
		return 0
	}
})

func allEq(ints []int) bool {
	val := ints[0]
	for _, i := range ints {
		if val != i {
			return false
		}
	}
	return true
}

func sum(ints []int) int {
	s := 0
	for _, i := range ints {
		s += i
	}
	return s
}

func sumIf(ints []int, f int) int {
	s := 0
	for _, i := range ints {
		if i == f {
			s += f
		}
	}
	return s
}

func allIncFrom(ints []int, f int) bool {
	curr := f
	for _, i := range ints {
		if i != curr {
			return false
		}
		curr++
	}
	return true
}

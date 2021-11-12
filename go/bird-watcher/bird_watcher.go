package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	s := 0
	for _, c := range birdsPerDay {
		s += c
	}
	return s
}

const daysInWeek = 7

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	l := daysInWeek * (week - 1)
	u := daysInWeek * week
	return TotalBirdCount(birdsPerDay[l:u])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, c := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = c + 1
		} else {
			birdsPerDay[i] = c
		}
	}
	return birdsPerDay
}

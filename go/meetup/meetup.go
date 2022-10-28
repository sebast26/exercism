package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

const week = 7

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	switch wSched {
	case First, Second, Third, Fourth:
		day := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
		day = day.AddDate(0, 0, week*int(wSched))
		daysToAdd := toAdd(day, wDay)
		return day.AddDate(0, 0, daysToAdd).Day()
	case Teenth:
		day := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
		daysToAdd := toAdd(day, wDay)
		return day.AddDate(0, 0, daysToAdd).Day()
	case Last:
		day := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC).Add(-1 * time.Minute)
		daysToSub := toSub(day, wDay)
		return day.AddDate(0, 0, -daysToSub).Day()
	}
	return 0
}

func toAdd(current time.Time, wDay time.Weekday) int {
	cDay := current.Weekday()
	if cDay > wDay {
		return week + int(wDay-cDay)
	}
	return int(wDay - cDay)
}

func toSub(current time.Time, wDay time.Weekday) int {
	cDay := current.Weekday()
	if cDay < wDay {
		return week - int(wDay-cDay)
	}
	return int(cDay - wDay)
}

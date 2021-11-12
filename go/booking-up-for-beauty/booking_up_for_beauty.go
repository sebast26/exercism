package booking

import "fmt"
import "time"

const Anniversary string = "September 15th 2012"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
  t, err := time.Parse("1/02/2006 15:04:05", date)
  if err != nil {
    panic(err)
  }
  return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
  t, err := time.Parse("January 2, 2006 15:04:05", date)
  if err != nil {
    panic(err)
  }
  return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
  // Thursday, July 25, 2019 13:45:00
  t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
  if err != nil {
    panic(err)
  }
  h := t.Hour()
  return h >= 12 && h < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
  t, err := time.Parse("1/2/2006 15:04:05", date)
  if err != nil {
    panic(err)
  }
  d := t.Format("Monday, January 2, 2006, at 15:04")
  return fmt.Sprintf("You have an appointment on %s.", d)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
  t, _ := time.Parse("January 2th 2006", Anniversary)
  thisYear, _, _ := time.Now().Date()
  _, anMonth, anDay := t.Date()
  return time.Date(thisYear, anMonth, anDay, 0, 0, 0, 0, time.UTC)
}

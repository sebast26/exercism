package twelve

import (
	"strings"
)

type day struct {
	num   string
	thing string
}

var days map[int]day

func init() {
	days = map[int]day{
		1:  {num: "first", thing: "a Partridge in a Pear Tree"},
		2:  {num: "second", thing: "two Turtle Doves"},
		3:  {num: "third", thing: "three French Hens"},
		4:  {num: "fourth", thing: "four Calling Birds"},
		5:  {num: "fifth", thing: "five Gold Rings"},
		6:  {num: "sixth", thing: "six Geese-a-Laying"},
		7:  {num: "seventh", thing: "seven Swans-a-Swimming"},
		8:  {num: "eighth", thing: "eight Maids-a-Milking"},
		9:  {num: "ninth", thing: "nine Ladies Dancing"},
		10: {num: "tenth", thing: "ten Lords-a-Leaping"},
		11: {num: "eleventh", thing: "eleven Pipers Piping"},
		12: {num: "twelfth", thing: "twelve Drummers Drumming"},
	}
}

func Verse(i int) string {
	var sb strings.Builder
	sb.WriteString("On the ")
	sb.WriteString(days[i].num)
	sb.WriteString(" day of Christmas my true love gave to me: ")

	var things []string
	for v := i; v > 1; v-- {
		things = append(things, days[v].thing)
	}
	sb.WriteString(strings.Join(things, ", "))

	if i > 1 {
		sb.WriteString(", and ")
	}
	sb.WriteString(days[1].thing)
	sb.WriteString(".")
	return sb.String()
}

func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n")
}

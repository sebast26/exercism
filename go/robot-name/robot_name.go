package robotname

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var names = map[string]bool{}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		gName := generateName()
		for ok := true; ok; _, ok = names[gName] {
			gName = generateName()
		}
		r.name = gName
		names[gName] = true
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func generateName() string {
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	sb.WriteRune('A' + rune(rand.Intn(26)))
	sb.WriteRune('A' + rune(rand.Intn(26)))
	sb.WriteString(strconv.Itoa(rand.Intn(10)))
	sb.WriteString(strconv.Itoa(rand.Intn(10)))
	sb.WriteString(strconv.Itoa(rand.Intn(10)))
	return sb.String()
}

package cui

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type color struct {
	Black       int
	White       int
	Red         int
	Purple      int
	Logo        int
	Yellow      int
	Green       int
	Menu        int
	Timestamp   int
	Header      int
	Notice      int
	Action      int
	QueryHeader int
}

var (
	Color = &color{
		Notice:      219,
		Action:      118,
		Black:       0,
		White:       15,
		Red:         160,
		Purple:      92,
		Logo:        75,
		Yellow:      11,
		Green:       119,
		Menu:        209,
		Timestamp:   247,
		Header:      57,
		QueryHeader: 11,
	}
)

// String returns a color escape string
func String(c int, str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m", c, str)
}

// StringFormat returns a color escape string with extra options
func StringFormat(c int, str string, args []string) string {
	return fmt.Sprintf("\x1b[38;5;%d;%sm%s\x1b[0m", c, strings.Join(args, ";"), str)
}

// StringFormatBoth fg and bg colors
// Thanks project komanda-cli
func StringFormatBoth(fg, bg int, str string, args []string) string {
	return fmt.Sprintf("\x1b[48;5;%dm\x1b[38;5;%d;%sm%s\x1b[0m", bg, fg, strings.Join(args, ";"), str)
}

// StringRandom returns a random colored string
func StringRandom(str string) string {
	return String(Random(22, 231), str)
}

// Random color number
func Random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

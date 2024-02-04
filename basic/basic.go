package basic

import (
	"encoding/binary"
	"fmt"
	"regexp"
	"time"
)

type Pattern string

const (
	KILOBYTE = 1024
	MEGABYTE = 1024 * KILOBYTE
	GIGABYTE = 1024 * MEGABYTE
	TERABYTE = 1024 * GIGABYTE
	PETABYTE = 1024 * TERABYTE
)

const (
	B   = 1
	KiB = B * 1024
	MiB = KiB * 1024
	GiB = MiB * 1024
	TiB = GiB * 1024
	PiB = TiB * 1024
	_   = PiB * 1024
)

const (
	Pattern01 Pattern = "#[st|nd|rd|th] [of] Jan[uary] [YY|YYYY]"
	Pattern02 Pattern = "Jan[uary] #[st|nd|rd|th] [in] [YY|YYYY]"
	Pattern03 Pattern = "Jan[uary] DD[,] [YY|YYYY]"
	Pattern04 Pattern = "Jan[uary] [YY|YYYY]"
	Pattern05 Pattern = "MM/DD/[YY|YYYY]"
	Pattern06 Pattern = "YYYY"
)

var (
	Months = map[string]time.Month{
		"JANUARY":   time.January,
		"FEBRUARY":  time.February,
		"MARCH":     time.March,
		"APRIL":     time.April,
		"MAY":       time.May,
		"JUNE":      time.June,
		"JULY":      time.July,
		"AUGUST":    time.August,
		"SEPTEMBER": time.September,
		"OCTOBER":   time.October,
		"NOVEMBER":  time.November,
		"DECEMBER":  time.December,

		"january":   time.January,
		"february":  time.February,
		"march":     time.March,
		"april":     time.April,
		"may":       time.May,
		"june":      time.June,
		"july":      time.July,
		"august":    time.August,
		"september": time.September,
		"october":   time.October,
		"november":  time.November,
		"december":  time.December,

		"jan": time.January,
		"feb": time.February,
		"mar": time.March,
		"apr": time.April,
		"jun": time.June,
		"jul": time.July,
		"aug": time.August,
		"sep": time.September,
		"oct": time.October,
		"nov": time.November,
		"dec": time.December,

		"JAN": time.January,
		"FEB": time.February,
		"MAR": time.March,
		"APR": time.April,
		"JUN": time.June,
		"JUL": time.July,
		"AUG": time.August,
		"SEP": time.September,
		"OCT": time.October,
		"NOV": time.November,
		"DEC": time.December,

		"01": time.January,
		"02": time.February,
		"03": time.March,
		"04": time.April,
		"05": time.May,
		"06": time.June,
		"07": time.July,
		"08": time.August,
		"09": time.September,
		"10": time.October,
		"11": time.November,
		"12": time.December,

		"1": time.January,
		"2": time.February,
		"3": time.March,
		"4": time.April,
		"5": time.May,
		"6": time.June,
		"7": time.July,
		"8": time.August,
		"9": time.September,
	}

	RegExDates = map[Pattern]*regexp.Regexp{
		Pattern01: regexp.MustCompile(`(?i)(\d{1,2})(st|nd|rd|th)?\s(?:of\s)?(January|Jan|February|Feb|March|Mar|April|Apr|May|June|Jun|July|Jul|August|Aug|September|Sep|October|Oct|November|Nov|December|Dec),?\s(\d{2,4})`),
		Pattern02: regexp.MustCompile(`(?i)(January|Jan|February|Feb|March|Mar|April|Apr|May|June|Jun|July|Jul|August|Aug|September|Sep|October|Oct|November|Nov|December|Dec)\s(\d{1,2})(st|nd|rd|th)?\s(?:in\s)?\s(\d{2,4})`),
		Pattern03: regexp.MustCompile(`(?i)(January|Jan|February|Feb|March|Mar|April|Apr|May|June|Jun|July|Jul|August|Aug|September|Sep|October|Oct|November|Nov|December|Dec)\s(\d{2})?,?\s(\d{4})`),
		Pattern04: regexp.MustCompile(`(?i)(January|Jan|February|Feb|March|Mar|April|Apr|May|June|Jun|July|Jul|August|Aug|September|Sep|October|Oct|November|Nov|December|Dec),?\s(\d{2,4})`),
		Pattern05: regexp.MustCompile(`(?i)(\d{1,2})\/(\d{1,2})\/(\d{2,4})`),
		Pattern06: regexp.MustCompile(`(\d{4})`),
	}
)

var NonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z]+`)
var NonNumericRegex = regexp.MustCompile(`[^0-9]+`)

// ReverseString reverses a string.
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReplaceNumbers(str string, repl string) string {
	return NonAlphanumericRegex.ReplaceAllString(str, repl)
}

func RemoveNumbers(str string) string {
	return ReplaceNumbers(str, "")
}

func ReplaceNonNumbers(str string, repl string) string {
	return NonNumericRegex.ReplaceAllString(str, repl)
}

func KeepNumbers(str string) string {
	return ReplaceNonNumbers(str, "")
}

func Uint64ToBytes(n uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)
	return b
}

func BytesToUint64(b []byte) (uint64, error) {
	if len(b) != 8 {
		return 0, fmt.Errorf("invalid slice length, expected len=8 got=%v", len(b))
	}
	return binary.BigEndian.Uint64(b), nil
}

package entry

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func ParseDate(date string) (time.Time, error) {
	return time.Parse(Timestamp, date)
}

func invalidDate(path string, err error) error {
	return fmt.Errorf("File has invalid date: %v\nERR: %v", path, err)
}

func MonthLang(t time.Time, lang string) string {
	if lang == "de" {
		return GermanMonths[t.Month()]
	}
	return t.Format("January")
}

var GermanMonths = map[time.Month]string{
	1:  "Januar",
	2:  "Februar",
	3:  "März",
	4:  "April",
	5:  "Mai",
	6:  "Juni",
	7:  "Juli",
	8:  "August",
	9:  "September",
	10: "Oktober",
	11: "November",
	12: "Dezember",
}

func Abbr(str string) string {
	a := ""
	i := 0
	for len(str) > 0 && i < 3 {
		r, size := utf8.DecodeRuneInString(str)
		a += string(r)

		str = str[size:]
		i++
	}
	return a
}

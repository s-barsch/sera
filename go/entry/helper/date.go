package helper

import (
	"fmt"
	p "path/filepath"
	"time"
	"unicode/utf8"
)

const Timestamp = "060102_150405"

func ParseId(id int64) time.Time {
	return time.Unix(id, 0)
}

func ToTimestamp(id int64) string {
	return ParseId(id).Format(Timestamp)
}

func ParseDatePath(path string) (time.Time, error) {
	id := Shorten(StripExt(p.Base(path)))

	date, err := ParseDate(id)
	if err != nil {
		return time.Time{}, DateErr(path, err)
	}

	return date, nil
}

func ParseDate(ts string) (time.Time, error) {
	return time.Parse(Timestamp, ts)
}

func DateErr(path string, err error) error {
	return fmt.Errorf("date error. %v. path: %v\n", err, path)
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

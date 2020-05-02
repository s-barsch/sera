package helper

import (
	"fmt"
	"strconv"
	"time"
)

func ShortenHash(a string) string {
	if len(a) > 3 {
		return a[len(a)-3:]
	}
	return a
}

var birth = int64(662774400)

func ToB16(t time.Time) string {
	return strconv.FormatInt(t.Unix()-birth, 16)
}

func ToB36(t time.Time) string {
	return strconv.FormatInt(t.Unix()-birth, 36)
}

func DecodeB16(s string) (string, error) {
	i, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return "", fmt.Errorf("decodeB16: %v", err)
	}
	t := time.Unix(i+birth, 0).UTC()
	return t.Format(Timestamp), nil
}

func EncodeAcronym(t time.Time) string {
	return ToB16(t)
}

func DecodeAcronym(acronym string) (string, error) {
	return DecodeB16(acronym)
}

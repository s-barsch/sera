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

const birth = int64(662774400)

func ToB16(t time.Time) string {
	return strconv.FormatInt(t.Unix()-birth, 16)
}

func ToB36(t time.Time) string {
	return strconv.FormatInt(t.Unix()-birth, 36)
}

func ParseHash(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("ParseB16: %v", err)
	}
	return i + birth, nil
	/*
		t := time.Unix(i+birth, 0).UTC()
		return t.Format(Timestamp), nil
	*/
}

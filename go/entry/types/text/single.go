package entry

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getSingleTextDate(path, dateEntry string) (time.Time, error) {
	date, err := ParseDate(Shorten(helper.StripExt(filepath.Base(path))))
	// Ignore error.
	if err == nil {
		return date, nil
	}

	date, err = ParseDate(dateEntry)
	if err != nil {
		err = fmt.Errorf("Cannot read date of %v\nErr: %v", path, err)
	}
	return date, err
}

func NewSingleText(path string, hold *Hold) (*Text, error) {
	file, err := NewFile(path, hold)
	if err != nil {
		return nil, err
	}

	parts, err := splitSingleText(path)
	if err != nil {
		return nil, err
	}

	info, err := unmarshalInfo([]byte(parts["info"]))
	if err != nil {
		return nil, fmt.Errorf("%v (%v)", err, path)
	}

	date, err := getSingleTextDate(path, info["date"])
	if err != nil {
		return nil, err
	}

	file.Id = date.Format(Timestamp)

	delete(parts, "info")

	style := "indent"
	if info["style"] != "" {
		style = strings.TrimSpace(info["style"])
	}

	html, err := markupLangs(parts, style)
	if err != nil {
		return nil, err
	}

	blank, err := stripLangs(html)
	if err != nil {
		return nil, err
	}

	text, err := hyphenateLangs(html)
	if err != nil {
		return nil, err
	}

	return &Text{
		File:  file,
		Info:  info,
		Date:  date,
		Text:  text,
		Blank: blank,
	}, nil
}

var partNames = map[int]string{
	0: "info",
	1: "de",
	2: "en",
}

func splitSingleText(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	parts := map[string]string{}

	s := bufio.NewScanner(io.Reader(f))
	buf := bytes.Buffer{}

	i := 0
	for s.Scan() {
		//fmt.Printf("SCAN: %v\n", s.Text())
		line := s.Text()
		if len(line) >= 3 && line[:3] == "---" {
			parts[partNames[i]] = buf.String()
			i++
			buf.Reset()
			continue
		}
		buf.WriteString(line)
		buf.WriteString("\n")
	}
	parts[partNames[i]] = buf.String()

	if i == 0 {
		parts["de"] = parts["info"]
		parts["info"] = ""
	}

	return parts, nil
}

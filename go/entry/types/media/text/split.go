package text

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"sacer/go/entry/helper"
)

var partNames = map[int]string{
	0: "info",
	1: "de",
	2: "en",
}

func splitTextFile(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, &helper.Err{
			Path: path,
			Func: "splitTextFile",
			Err:  err,
		}
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

package text

/*
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

}
*/

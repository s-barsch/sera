package el

import (
	"io/ioutil"
)

func NewMultiText(path string, hold *Hold) (*Text, error) {
	file, err := NewFile(path, hold)
	if err != nil {
		return nil, err
	}

	date, err := getFilenameDate(path)
	if err != nil {
		return nil, err
	}

	info := map[string]string{}

	langs, err := ReadMultiText(path)
	if err != nil {
		return nil, err
	}

	// TODO: Undetermined paragraph style.
	html, err := markupLangs(langs, "lines")
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
		File: file,

		Date: date,
		Info: info,

		Text:  text,
		Blank: blank,
	}, nil
}

func ReadMultiText(path string) (map[string]string, error) {
	de, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	en, err := ioutil.ReadFile(enFile(path))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"de": string(de),
		"en": string(en),
	}, nil
}

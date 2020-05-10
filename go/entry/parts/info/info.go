package info

import (
	//"fmt"
	"io"
	"os"
	"stferal/go/entry/helper"
	//"stferal/go/entry/helper/hyph"

	"gopkg.in/yaml.v2"

	p "path/filepath"
	"strings"
)

type Info map[string]string

func ReadDirInfo(path string) (Info, error) {
	return ParseInfoFile(p.Join(path, "info"))
}

func ReadFileInfo(path string) (Info, error) {
	return ParseInfoFile(fileInfo(path))
}

func fileInfo(path string) string {
	return path + ".info"
}

func HasFileInfo(path string) bool {
	_, err := os.Stat(fileInfo(path))
	if err == nil {
		return true
	}
	return false
}

func ParseInfoFile(path string) (Info, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "ParseInfoFile",
	}

	i := map[string]string{}

	f, err := os.Open(path)
	if err != nil {
		fnErr.Err = err
		return i, err
	}
	defer f.Close()

	d := yaml.NewDecoder(io.Reader(f))
	err = d.Decode(&i)
	if err != nil {
		fnErr.Err = err
		return i, err
	}

	for k, v := range i {
		delete(i, k)
		i[norm(k)] = trim(v)
	}

	// TODO: do this later
	// err = i.hyphenateText()

	return i, nil

}

func norm(str string) string {
	return helper.Normalize(str)
}

func trim(str string) string {
	return strings.TrimSpace(str)
}

func (i Info) Title(lang string) string {
	return i.Field("title", lang)
}

func (i Info) HyphTitle(lang string) string {
	return i.Field("title-hyph", lang)
}

/*
func (i Info) TitleUpper(lang string) string {
	title := i.Field("title-hyph", lang)
	return s.Replace(title, "ß", "ss", -1)
}
*/

func (i Info) Caption(lang string) string {
	return i.Field("caption", lang)
}

func (i Info) Description(lang string) string {
	return i.Field("description", lang)
}

func (i Info) Alt(lang string) string {
	return i.Field("alt", lang)
}

func (i Info) Slug(lang string) string {
	return i.Field("slug", lang)
}

/*
func (i Info) Label(lang string) string {
	if label := i.Field("label", lang); label != "" {
		return label
	}
	return i.Title(lang)
}
*/

func (i Info) Field(key, lang string) string {
	if lang != "de" {
		return i[key+"-"+lang]
	}
	return i[key]
}

/*
func (i Info) Location() string {
	return i["location"]
}
*/

/*
func slug(f *File, info Info, lang string) string {
	if lang == "de" {
		return helper.StripExt(f.Base())
	}
	if slug := info.Field("slug", lang); slug != "" {
		return slug
	}
	return slugify(info.Title(lang))
}
*/

func UnmarshalInfo(input []byte) (Info, error) {
	i := map[string]string{}
	err := yaml.Unmarshal(input, &i)
	if err != nil {
		return i, err
	}

	for k, v := range i {
		delete(i, k)
		i[norm(k)] = trim(v)
	}

	return i, nil
}

/*
var hyphDirections = map[string]string{
	"caption":       "de",
	"caption-en":    "en",
	"alt":           "de",
	"alt-en":        "en",
	"transcript":    "de",
	"transcript-en": "en",
}

func (i Info) hyphenateText() error {
	for k, l := range hyphDirections {
		if i[k] != "" {
			b, err := hyph.HyphenateText(i[k], l)
			if err != nil {
				return err
			}
			i[k] = string(b)
		}
	}
	s, err := hyph.HyphenateText(i["title"], "de")
	if err != nil {
		return err
	}
	i["title-hyph"] = s
	s, err = hyph.HyphenateText(i["title-en"], "en")
	if err != nil {
		return err
	}
	i["title-hyph-en"] = s
	return nil
}
*/

/*
func (i Info) State() string {
	state := i["state"]

	if state == "" {
		return "live"
	}

	return state
}
*/

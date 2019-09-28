package el

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"stferal/pkg/el/hyph"
	//"io/ioutil"
	"path/filepath"
	"strings"
)

type Info map[string]string

func ReadInfo(path string) (Info, error) {
	return parseInfoFile(infoPath(path))
}

func parseInfoFile(path string) (Info, error) {
	//i := map[string]string{}
	i := Info{}
	f, err := os.Open(path)
	if err != nil {
		return i, fmt.Errorf("parseInfoFile: %v", err)
	}
	defer f.Close()

	d := yaml.NewDecoder(io.Reader(f))
	err = d.Decode(&i)
	if err != nil {
		return i, fmt.Errorf("parseInfoFile: %v, %v", err, path)
	}

	for k, v := range i {
		delete(i, k)
		i[strings.ToLower(k)] = v
	}

	err = i.hyphenateText()

	return i, nil

}

func (i Info) State() string {
	state := i["state"]

	if state == "" {
		return "live"
	}

	return state
}

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

func (i Info) Title(lang string) string {
	return i.Field("title", lang)
}

func (i Info) TitleUpper(lang string) string {
	title := i.Field("title-hyph", lang)
	return strings.Replace(title, "ß", "ss", -1)
}

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
	if slug := i.Field("slug", lang); slug != "" {
		return slug
	}
	return Normalize(i.Title(lang))
}

func (i Info) Label(lang string) string {
	if label := i.Field("label", lang); label != "" {
		return label
	}
	return i.Title(lang)
}

func (i Info) Location() string {
	return i["location"]
}

func (i Info) Field(key, lang string) string {
	if lang != "de" {
		return i[key+"-"+lang]
	}
	return i[key]
}

/*
func slug(f *File, info Info, lang string) string {
	if lang == "de" {
		return stripExt(f.Base())
	}
	if slug := info.Field("slug", lang); slug != "" {
		return slug
	}
	return slugify(info.Title(lang))
}
*/

func stripExt(base string) string {
	i := strings.LastIndex(base, ".")
	if i <= 0 {
		return base
	}
	return base[:i]
}

func infoPath(path string) string {
	return filepath.Join(path, "info")
}

func unmarshalInfo(input []byte) (Info, error) {
	i := map[string]string{}
	err := yaml.Unmarshal(input, &i)
	if err != nil {
		return i, err
	}

	for k, v := range i {
		delete(i, k)
		i[strings.ToLower(k)] = v
	}

	return i, nil
}

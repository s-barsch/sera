package info

import (
	"io"
	"os"
	p "path/filepath"
	"strings"

	"g.rg-s.com/sferal/go/entry/tools"

	"gopkg.in/yaml.v3"
)

func (i Info) Copy() Info {
	m := map[string]string{}

	for k, v := range i {
		m[k] = v
	}
	return m
}

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
	return err == nil
}

func ParseInfoFile(path string) (Info, error) {
	fnErr := &tools.Err{
		Path: path,
		Func: "ParseInfoFile",
	}

	i := Info{}

	f, err := os.Open(path)
	if err != nil {
		fnErr.Err = err
		return i, fnErr
	}
	defer f.Close()

	d := yaml.NewDecoder(io.Reader(f))
	err = d.Decode(&i)
	if err != nil {
		fnErr.Err = err
		return i, fnErr
	}

	i.Clean()
	i.Hyphenate()

	return i, nil

}

func UnmarshalInfo(input []byte) (Info, error) {
	i := Info{}
	err := yaml.Unmarshal(input, &i)
	if err != nil {
		return i, err
	}

	i.Clean()
	i.Hyphenate()

	return i, nil
}

func (i Info) Clean() {
	for k, v := range i {
		delete(i, k)
		key := tools.Normalize(k)

		// Keep new lines EOF
		if name(key) == "transcript" {
			i[key] = v
			continue
		}

		i[key] = strings.TrimSpace(v)
	}
}

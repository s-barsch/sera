package info

import (
	"gopkg.in/yaml.v2"
	"io"
	"os"
	p "path/filepath"
	"sacer/go/entry/tools"
	"strings"
)

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



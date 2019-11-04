package paths

import (
	"stferal/pkg/entry"
	"strings"
)

type Path struct {
	Page       string
	Holds      []string
	Name       string
	Acronym    string
	Type       string
	Descriptor string
}

// files (punkt)
// cache (punkt)
// hold mit id
// hold ohne id
// el mit id

func isFile(path string) bool {
	for _, c := range path {
		switch c {
		case '.':
			return true
		case '/':
			return false
		}
	}
	return false
}

func last(chain []string) string {
	if len(chain) == 0 {
		return ""
	}
	return chain[len(chain)-1]
}

func removeLast(chain []string) []string {
	if len(chain) == 0 {
		return chain
	}
	return chain[:len(chain)-1]
}

func splitName(str string) (name, acronym string) {
	//TODO: Function doesn’t discern between “innen-aussen” and “form-a9faad9”.
	i := strings.LastIndex(str, "-")
	if i < 0 {
		_, err := entry.DecodeAcronym(str)
		if err != nil {
			return str, ""
		}
		return "", str
	}
	// /index/art-34188329/

	acronym = str[i+1:]

	_, err := entry.DecodeAcronym(acronym)
	if err != nil {
		return str, ""
	}

	return str[:i], acronym
}

func decodeDirName(str string) (name, acronym string) {
	// case: /index/34188329 or /index/art/
	_, err := entry.DecodeAcronym(str)
	if err == nil {
		return "", str
	}
	return str, ""
}

func Split(path string) *Path {
	chain := strings.Split(strings.Trim(path, "/"), "/")

	page := chain[0]
	chain = chain[1:]

	typ := ""
	descriptor := ""

	for i, c := range chain {
		if c == "files" || c == "cache" {
			typ = c
			descriptor = strings.Join(chain[i+1:], "/")
			chain = chain[:i]
			break
		}
	}

	name, acronym := splitName(last(chain))

	holds := removeLast(chain)

	return &Path{
		Page:       page,
		Holds:      holds,
		Name:       name,
		Acronym:    acronym,
		Type:       typ,
		Descriptor: descriptor,
	}
}

// descriptor == 160403_124512-1600.jpg
//				 |             |
//				 filename	   size
func SplitDescriptor(desc string) (filename, size string) {
	i := strings.LastIndex(desc, "-")
	if i < 0 {
		return desc, ""
	}
	j := strings.LastIndex(desc, ".")
	if j < 0 {
		// Descriptor "file_no_ext-1600", kommt eigentlich nicht vor.
		return desc[:i], desc[i+1:]
	}
	filename = desc[:i] + desc[j:]
	size = desc[i+1 : j]
	return
}

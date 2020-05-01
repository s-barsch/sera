package paths

import (
	"stferal/go/entry/helper"
	"strings"
)

type Path struct {
	Page    string
	Holds   []string
	Name    string
	Acronym string
	Subdir  string
	Subpath string
}

/*
	/graph/2020/03/09-36e55605/cache/200310_012140-1280.jpg

    &paths.Path{
            Page:       "graph",
            Holds:      {"2020", "03"},
            Name:       "09",
            Acronym:    "36e55605",
            Subdir:     "cache",
            Subpath: "200310_012140-1280.jpg",
        }
*/

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
	i := strings.LastIndex(str, "-")
	if i < 0 {
		_, err := helper.DecodeAcronym(str)
		if err != nil {
			return str, ""
		}
		return "", str
	}

	acronym = str[i+1:]

	_, err := helper.DecodeAcronym(acronym)
	if err != nil {
		return str, ""
	}

	return str[:i], acronym
}

func decodeDirName(str string) (name, acronym string) {
	_, err := helper.DecodeAcronym(str)
	if err == nil {
		return "", str
	}
	return str, ""
}

func Split(path string) *Path {
	chain := strings.Split(strings.Trim(path, "/"), "/")

	page := chain[0]
	chain = chain[1:]

	subdir := ""
	subpath := ""

	for i, c := range chain {
		if c == "files" || c == "cache" {
			subdir = c
			subpath = strings.Join(chain[i+1:], "/")
			chain = chain[:i]
			break
		}
	}

	name, acronym := splitName(last(chain))

	holds := removeLast(chain)

	return &Path{
		Page:    page,
		Holds:   holds,
		Name:    name,
		Acronym: acronym,
		Subdir:  subdir,
		Subpath: subpath,
	}
}

// subpath == 160403_124512-1600.jpg
//			  |             |
//			  filename	    size

func SplitSubpath(subp string) (filename, size string) {
	i := strings.LastIndex(subp, "-")
	if i < 0 {
		return subp, ""
	}
	// expect a file extension
	j := strings.LastIndex(subp, ".")
	if j < 0 {
		return subp[:i], subp[i+1:]
	}
	filename = subp[:i] + subp[j:]
	size = subp[i+1 : j]
	return
}

package paths

import (
	"strings"
)

type Path struct {
	Page    string
	Holds   []string
	Name    string
	Hash    string
	Subdir  string
	Subpath string
}

/*
	/graph/2020/03/09-36e55605/cache/200310_012140-1280.jpg

    &paths.Path{
            Page:       "graph",
            Holds:      {"2020", "03"},
            Name:       "09",
            hash:    "36e55605",
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

func splitName(str string) (name, hash string) {
	i := strings.LastIndex(str, "-")
	if i < 0 {
		return "", str
	}
	return str[:i], str[i+1:]
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

	name, hash := splitName(last(chain))

	holds := removeLast(chain)

	return &Path{
		Page:    page,
		Holds:   holds,
		Name:    name,
		Hash:    hash,
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

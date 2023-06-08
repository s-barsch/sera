package paths

import (
	"regexp"
	"sacer/go/entry/tools"
	"strings"
)

type Path struct {
	Raw     string
	Chain   []string
	Slug    string
	Hash    string
	SubDir  string
	SubFile *SubFile
}

type SubFile struct {
	Name, Size string
}

func (p *Path) Lang() string {
	if len(p.Chain) > 0 {
		return p.Chain[0]
	}
	return ""
}

func (p *Path) Section() string {
	if len(p.Chain) > 1 {
		section := p.Chain[1]
		if section == "cine" {
			return "kine"
		}
		return section
	}
	return ""
}

func (p *Path) IsFile() bool {
	if p.SubDir != "" {
		return true
	}
	return false
}

// || strings.Contains(p.Raw, ".") {

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

func Split(path string) *Path {
	chain := strings.Split(strings.Trim(path, "/"), "/")

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

	slug, hash := splitName(last(chain))

	chain = removeLast(chain)

	return &Path{
		Raw:     path,
		Chain:   chain,
		Slug:    slug,
		Hash:    hash,
		SubDir:  subdir,
		SubFile: SplitSubpath(subpath),
	}
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

func IsMergedMonths(str string) bool {
	return regexp.MustCompile("\\d{2}-\\d{2}").MatchString(str)
}

func splitName(str string) (slug, hash string) {
	i := strings.LastIndex(str, "-")
	if i < 0 {
		return discernName(str)
	}
	// for merged months "11-12"
	if i == 2 && IsMergedMonths(str) {
		return str, ""
	}
	return str[:i], str[i+1:]
}

func discernName(str string) (slug, hash string) {
	// for year pages /graph/2006
	if len(str) < 5 {
		return str, ""
	}
	_, err := tools.ParseHash(str)
	if err == nil {
		return "", str
	}
	return str, ""
}

// subpath == 160403_124512-1600.jpg
//			  |             |
//			  filename	    size

// 160403_124512-1600.jpg -> (160403_124512.jpg) (1600)
func SplitSubpath(subp string) *SubFile {
	i := strings.LastIndex(subp, "-")
	if i < 0 {
		return &SubFile{
			Name: subp,
		}
	}
	j := strings.LastIndex(subp, ".")
	if j < 0 {
		// No file extension.
		// 160403_124512-1600 -> (160403_124512) (1600)
		return &SubFile{
			Name: subp[:i],
			Size: subp[i+1:],
		}
	}
	// Remove size and put filename back together.
	return &SubFile{
		Name: subp[:i] + subp[j:],
		Size: subp[i+1 : j],
	}
}

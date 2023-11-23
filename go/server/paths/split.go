package paths

import (
	"regexp"
	"sacer/go/entry/tools"
	"strings"
)

type Path struct {
	Raw    string
	Chain  []string
	Slug   string
	Hash   string
	Folder string
	File   *File
}

type File struct {
	Name, Option, Ext string
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
		return section
	}
	return ""
}

func (p *Path) IsFile() bool {
	if p.Folder != "" {
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

	folder := ""
	subpath := ""

	for i, c := range chain {
		if c == "files" || c == "cache" {
			folder = c
			subpath = strings.Join(chain[i+1:], "/")
			chain = chain[:i]
			break
		}
	}

	slug, hash := splitName(last(chain))

	chain = removeLast(chain)

	return &Path{
		Raw:    path,
		Chain:  chain,
		Slug:   slug,
		Hash:   hash,
		Folder: folder,
		File:   SplitFile(subpath),
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

// filepath == 160403_124512-1600.jpg
//			  |             |
//			  filename	    size

// 160403_124512-1600.jpg -> (160403_124512.jpg) (1600)
func SplitFile(filep string) *File {
	i := strings.LastIndex(filep, "-")
	if i < 0 {
		return &File{
			Name: filep,
		}
	}
	j := strings.LastIndex(filep, ".")
	if j < 0 {
		// No file extension.
		// 160403_124512-1600 -> (160403_124512) (1600)
		return &File{
			Name:   filep[:i],
			Option: filep[i+1:],
		}
	}
	// Remove size and put filename back together.
	return &File{
		Name:   filep[:i] + filep[j:],
		Option: filep[i+1 : j],
		Ext:    filep[j+1:],
	}
}

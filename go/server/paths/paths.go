package paths

import (
	"log"
	"regexp"
	"strings"

	"g.rg-s.com/sera/go/entry/tools"
)

type Path struct {
	Path   string
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
	return p.File != nil
}

func Split(path string) *Path {
	rawChain := strings.Split(strings.Trim(path, "/"), "/")
	chain, folder, subpath := ExtractFolder(rawChain)
	cutChain, name := ExtractName(chain)

	slug, hash := splitSlugHash(name)

	split, err := SplitFile(subpath)
	if err != nil {
		log.Println(err)
		split = &File{Name: subpath}
	}

	return &Path{
		Path:   path,
		Chain:  cutChain,
		Slug:   slug,
		Hash:   hash,
		Folder: folder,
		File:   split,
	}
}

func ExtractName(chain []string) (cutChain []string, name string) {
	return removeLast(chain), last(chain)
}

func ExtractFolder(chain []string) (cut []string, folder, subpath string) {
	for i, c := range chain {
		if i > 1 && (c == "files" || c == "img") {
			return chain[:i], c, strings.Join(chain[i+1:], "/")
		}
	}
	return chain, folder, subpath
}

// Slug and hash are seperated by a dash
// lonely-3f397f82
// |slug  |hash
func splitSlugHash(str string) (slug, hash string) {
	i := strings.LastIndex(str, "-")
	// if no dash present: further checks
	if i < 0 {
		return discernName(str)
	}
	// for merged months "11-12": check via regex
	if i == 2 && IsMergedMonths(str) {
		return str, ""
	}
	// if proper split present: just return
	return str[:i], str[i+1:]
}

// "name" understood as a mixture of slug and hash or
// one of the two.
func discernName(str string) (slug, hash string) {
	// for year pages /graph/2006: all slug
	if len(str) < 5 {
		return str, ""
	}
	// if parses: must be all hash
	_, err := tools.ParseHash(str)
	if err == nil {
		return "", str
	}
	return str, ""
}

func IsMergedMonths(str string) bool {
	return regexp.MustCompile(`\d{2}-\d{2}`).MatchString(str)
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

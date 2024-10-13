package paths

import (
	"log"
	"regexp"
	"strings"

	"g.rg-s.com/sera/go/entry/tools"
)

type Split struct {
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

func (p *Split) Lang() string {
	if len(p.Chain) > 0 {
		return p.Chain[0]
	}
	return ""
}

func (p *Split) Section() string {
	if len(p.Chain) > 1 {
		section := p.Chain[1]
		return section
	}
	return ""
}

func (p *Split) IsFile() bool {
	return p.File != nil
}

// /en/cache/24-08/ 10-super-theory- 3f412b02 /img/      cover-480.webp"
// (chain)          (slug)           (hash)   (folder)   (file)
func SplitPath(path string) *Split {
	rawChain := strings.Split(strings.Trim(path, "/"), "/")

	cutChain, folder, split := extractSplitFile(rawChain)

	chain, slug, hash := extractSlugHash(cutChain)

	return &Split{
		Path:   path,
		Chain:  chain,
		Slug:   slug,
		Hash:   hash,
		Folder: folder,
		File:   split,
	}
}

func extractSlugHash(chain []string) ([]string, string, string) {
	slug, hash := splitSlugHash(last(chain))
	return removeLast(chain), slug, hash
}

func extractSplitFile(rawChain []string) (chain []string, folder string, file *File) {
	chain, folder, subpath := extractFolder(rawChain)

	split, err := SplitFile(subpath)
	if err != nil {
		log.Println(err)
		split = &File{Name: subpath}
	}

	return chain, folder, split
}

func extractFolder(chain []string) (cut []string, folder, subpath string) {
	for i, c := range chain {
		if i > 1 && (c == "files" || c == "img") {
			return chain[:i], c, strings.Join(chain[i+1:], "/")
		}
	}
	return chain, folder, subpath
}

// Slug and hash are seperated by a dash: lonely-3f397f82
func splitSlugHash(str string) (slug, hash string) {
	i := strings.LastIndex(str, "-")
	// if contains dash: check if it's "11-12" format
	if i == 2 && IsMergedMonths(str) {
		return str, ""
	}
	// if contains dash: check if last word is in fact hash
	if i > 0 {
		slug = str[:i]
		hash = str[i+1:]
		if tools.IsHash(hash) {
			return slug, hash
		}
	}
	// short values like years and months are no hash
	if len(str) < 5 {
		return str, ""
	}
	// check if it is all hash
	if tools.IsHash(str) {
		return "", str
	}
	// must be all slug
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

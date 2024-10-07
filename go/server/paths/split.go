package paths

import (
	"fmt"
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
	return p.Folder != ""
}

func Split(path string) *Path {
	chain := strings.Split(strings.Trim(path, "/"), "/")

	folder := ""
	subpath := ""

	for i, c := range chain {
		if i > 1 && (c == "files" || c == "img") {
			folder = c
			subpath = strings.Join(chain[i+1:], "/")
			chain = chain[:i]
			break
		}
	}

	slug, hash := splitName(last(chain))

	chain = removeLast(chain)

	split, err := SplitFile(subpath)
	if err != nil {
		log.Println(err)
		split = &File{
			Name: subpath,
		}
	}

	return &Path{
		Path:   path,
		Chain:  chain,
		Slug:   slug,
		Hash:   hash,
		Folder: folder,
		File:   split,
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
	return regexp.MustCompile(`\d{2}-\d{2}`).MatchString(str)
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
//			  name			option
//
// New filename: 160403_124512-1600.jpg
// Option: 1600
// Ext: "jpg"

func SplitFile(filep string) (*File, error) {
	if len(filep) >= len("vtt/x.de.vtt") && filep[:3] == "vtt" {
		return splitVTT(filep)
	}
	return splitMedia(filep)
}

// splitMedia: for images and video

func splitMedia(filep string) (*File, error) {
	i := strings.LastIndex(filep, "-")
	j := strings.LastIndex(filep, ".")

	return splitFileParameters(filep, i, j)
}

// Sample filepath: vtt/x.de.vtt

func splitVTT(filep string) (*File, error) {
	i := strings.Index(filep, ".")
	j := strings.LastIndex(filep, ".")

	return splitFileParameters(filep, i, j)
}

// 160403_124512-1600.jpg -> (160403_124512.jpg) (1600)

func splitFileParameters(filep string, i, j int) (*File, error) {
	if i < 0 || j < 0 || i == j {
		return nil, fmt.Errorf("splitFileParameters: errornous filename: %v", filep)
	}

	return &File{
		Name:   filep[:i] + filep[j:],
		Option: filep[i+1 : j],
		Ext:    filep[j+1:],
	}, nil
}

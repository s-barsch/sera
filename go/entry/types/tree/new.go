package tree

import (
	p "path/filepath"
	"time"

	"g.rg-s.com/sferal/go/entry"
	"g.rg-s.com/sferal/go/entry/file"
	"g.rg-s.com/sferal/go/entry/info"
	"g.rg-s.com/sferal/go/entry/tools"
	"g.rg-s.com/sferal/go/entry/tools/script"
	"g.rg-s.com/sferal/go/entry/types/image"
	"g.rg-s.com/sferal/go/entry/types/set"
)

type Tree struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	entries entry.Entries
	Trees   Trees

	Cover *image.Image

	Footnotes script.Footnotes

	Summary        *script.Script
	SummaryPrivate *script.Script
}

func (t *Tree) Copy() *Tree {
	return &Tree{
		parent: t.parent,
		file:   t.file.Copy(),

		date: t.date,
		info: t.info.Copy(),

		entries: t.entries,
		Trees:   t.Trees,

		Cover: t.Cover,

		Footnotes: t.Footnotes,

		Summary:        t.Summary,
		SummaryPrivate: t.SummaryPrivate,
	}
}

type Trees []*Tree

func ReadTree(path string, parent *Tree) (*Tree, error) {
	fnErr := &tools.Err{
		Path: path,
		Func: "ReadTree",
	}

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	inf, err := readTreeInfo(path, parent)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	date, err := tools.ParseTimestamp(inf["date"])
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s := &Tree{
		parent: parent,
		file:   file,

		date: date,
		info: inf,
	}

	trees, err := readTrees(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	entries, err := readEntries(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	cover := &image.Image{}

	cover, entries = extractCover(entries)

	summary := getScript(inf, "summary")
	summaryPrivate := getScript(inf, "summary-private")

	s.Summary = summary
	s.SummaryPrivate = summaryPrivate

	s.entries = entries
	s.Trees = trees
	s.Cover = cover

	s.Footnotes = set.NumberFootnotes(s.Entries())

	return s, nil
}

func readTreeInfo(path string, parent *Tree) (info.Info, error) {
	if !isGraph(path, parent) {
		return info.ReadDirInfo(path)
	}
	return readGraphInfo(path, parent)
}

// Function only needed here, not in readTrees or readEntries.
// Because in these, #parent# will always be defined.
func isGraph(path string, parent *Tree) bool {
	if parent == nil {
		return isGraphSection(p.Base(path))
	}
	return isGraphSection(parent.Section())
}

func isGraphSection(section string) bool {
	switch section {
	case "graph", "cache":
		return true
	}
	return false
}

func getScript(i info.Info, key string) *script.Script {
	langs := extractScript(i, key)
	if langs == nil {
		return nil
	}
	script := script.RenderScript(langs)
	script.NumberFootnotes(1)

	return script
}

func extractScript(i info.Info, key string) script.LangMap {
	langs := script.LangMap{}
	for l := range tools.Langs {
		if l != "de" {
			key += "-" + l
		}
		langs[l] = i[key]
	}
	if langs["de"] == "" && langs["en"] == "" {
		return nil
	}
	return langs
}

func extractCover(es entry.Entries) (*image.Image, entry.Entries) {
	for i, e := range es {
		if e.File().Name() == "cover.jpg" {
			img, ok := e.(*image.Image)
			if ok {
				return img, append(es[:i], es[i+1:]...)
			}
		}
	}
	return nil, es
}

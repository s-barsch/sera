package extra

import (
	"fmt"
	"log"
	"net/http"
	p "path/filepath"
	"strings"
	"time"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/tools"
	"g.rg-s.com/sera/go/entry/types/set"
	"g.rg-s.com/sera/go/entry/types/tree"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/paths"
	"g.rg-s.com/sera/go/viewer"
)

/*

	TODO: This package should be simplified.

*/

func ServeFile(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := serveFile(w, r, v, m)
		if err != nil {
			log.Println(err)
			http.NotFound(w, r)
		}
	}
}

func serveFile(w http.ResponseWriter, r *http.Request, v *viewer.Viewer, m *meta.Meta) error {
	section := m.Split.Section()

	tree, ok := v.Store.TreeByString(section)
	if !ok {
		return fmt.Errorf("no tree found for section %q", section)
	}

	e, err := getEntry(tree, m.Split)
	if err != nil {
		return err
	}

	col, ok := e.(entry.Collection)

	if !ok {
		return serveSingleBlob(w, r, m, e)
	}

	return serveCollectionBlob(w, r, m, col)
}

func serveSingleBlob(w http.ResponseWriter, r *http.Request, m *meta.Meta, e entry.Entry) error {
	blob, ok := e.(entry.Blob)
	if !ok {
		return fmt.Errorf("file to serve (%v) is no blob", e.File().Name())
	}

	location, err := blob.Location(m.Split.File.Ext, m.Split.File.Option)
	if err != nil {
		return err
	}
	serveStatic(w, r, location)
	return nil
}

func serveCollectionBlob(w http.ResponseWriter, r *http.Request, m *meta.Meta, col entry.Collection) error {
	name := baseName(m.Split.File.Name)
	for _, e := range col.Entries() {
		if baseName(e.File().Name()) == name {
			return serveSingleBlob(w, r, m, e)
		}
	}

	if name := m.Split.File.Name; len(name) > 5 && name[:5] == "cover" {
		set, ok := col.(*set.Set)
		if ok && set.Cover != nil {
			return serveSingleBlob(w, r, m, set.Cover)
		}
		t, ok := col.(*tree.Tree)
		if ok && t.Cover != nil {
			return serveSingleBlob(w, r, m, t.Cover)
		}
		return fmt.Errorf("serveCollectionBlob: Cover %v not found", m.Split.File.Name)
	}

	return fmt.Errorf("serveCollectionBlob: File %v not found", m.Split.File.Name)
}

func baseName(name string) string {
	name = stripBlur(name)
	return stripSize(name)
}

func stripSize(name string) string {
	i := strings.LastIndex(name, "-")
	if i > 0 {
		return name[:i]
	}
	return name
}

func stripBlur(name string) string {
	name = tools.StripExt(p.Base(name))
	if l := len(name); l > 4 && name[l-4] == '_' {
		return name[:l-4]
	}
	return name
}

func getEntry(t *tree.Tree, path *paths.Split) (entry.Entry, error) {
	hash := path.Hash
	if hash == "" {
		h, err := getMonthHash(path)
		if err != nil {
			return nil, err
		}
		hash = h
	}
	return t.LookupEntryHash(hash)
}

func getMonthHash(path *paths.Split) (string, error) {
	if len(path.Chain) != 3 {
		return "", fmt.Errorf("getMonthEntry: wrong month format. %v", path.Raw)
	}

	slug := path.Slug
	if paths.IsMergedMonths(path.Slug) {
		slug = slug[:2]
	}

	date, err := time.Parse("200601--150405", path.Chain[2]+slug+"--000001")
	if err != nil {
		return "", err
	}

	return tools.ToB16(date), nil
}

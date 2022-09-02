package extra

import (
	"fmt"
	"log"
	"net/http"
	p "path/filepath"
	"time"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/set"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/meta"
	"sacer/go/server/paths"
	"strings"
)

/*

	TODO: This package should be simplified.

*/

func ServeFile(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta, path *paths.Path) {
	err := serveFile(s, w, r, m, path)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
	}
}

func serveFile(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta, path *paths.Path) error {
	section := path.Section()
	tree := s.Trees[section].Access(m.Auth.Subscriber)[m.Lang]

	e, err := getEntry(tree, path)
	if err != nil {
		return err
	}

	col, ok := e.(entry.Collection)

	if !ok {
		return serveSingleBlob(w, r, e, path)
	}

	return serveCollectionBlob(w, r, col, path)
}

func serveSingleBlob(w http.ResponseWriter, r *http.Request, e entry.Entry, path *paths.Path) error {
	blob, ok := e.(entry.Blob)
	if !ok {
		return fmt.Errorf("File to serve (%v) is no blob.", e.File().Name())
	}

	m, ok := e.(entry.Media)
	if ok {
		switch p.Ext(path.SubFile.Name) {
		case ".vtt":
			serveStatic(w, r, m.CaptionLocation(path.SubFile.Size))
			return nil
		}
	}
	location, err := blob.Location(path.SubFile.Size)
	if err != nil {
		return err
	}
	serveStatic(w, r, location)
	return nil
}

func serveCollectionBlob(w http.ResponseWriter, r *http.Request, col entry.Collection, path *paths.Path) error {
	name := baseName(path.SubFile.Name)
	for _, e := range col.Entries() {
		if baseName(e.File().Name()) == name {
			return serveSingleBlob(w, r, e, path)
		}
	}

	if name := path.SubFile.Name; len(name) > 5 && name[:5] == "cover" {
		set, ok := col.(*set.Set)
		if ok && set.Cover != nil {
			return serveSingleBlob(w, r, set.Cover, path)
		}
		t, ok := col.(*tree.Tree)
		if ok && t.Cover != nil {
			return serveSingleBlob(w, r, t.Cover, path)
		}
		return fmt.Errorf("serveCollectionBlob: Cover %v not found.", path.SubFile.Name)
	}

	return fmt.Errorf("serveCol !N!,onBlob: File %v not found.", path.SubFile.Name)
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

func getEntry(t *tree.Tree, path *paths.Path) (entry.Entry, error) {
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

func getMonthHash(path *paths.Path) (string, error) {
	if len(path.Chain) != 2 {
		return "", fmt.Errorf("getMonthEntry: wrong month format. %v", path.Raw)
	} 

	date, err := time.Parse("200601--150405", path.Chain[1] + path.Slug + "--000001")
	if err != nil {
		return "", err
	}

	return tools.ToB16(date), nil
}






package extra

import (
	"fmt"
	"log"
	"net/http"

	//"path/filepath"
	"stferal/go/entry"
	"stferal/go/paths"
	"stferal/go/server"
)

func ServeFile(s *server.Server, w http.ResponseWriter, r *http.Request, path *paths.Path) {
	err := serveFile(s, w, r, path)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
	}
}

func serveFile(s *server.Server, w http.ResponseWriter, r *http.Request, path *paths.Path) error {
	e, err := s.Trees[path.Section()].LookupEntryHash(path.Hash)
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
	serveStatic(w, r, blob.Location(path.SubFile.Size))
	return nil
}

func serveCollectionBlob(w http.ResponseWriter, r *http.Request, col entry.Collection, path *paths.Path) error {
	for _, e := range col.Entries() {
		if e.File().Name() == path.SubFile.Name {
			return serveSingleBlob(w, r, e, path)
		}
	}
	return fmt.Errorf("serveCollectionBlob: File %v not found.", path.SubFile.Name)
}

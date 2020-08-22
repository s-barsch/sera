package extra

import (
	"fmt"
	"log"
	"net/http"

	"stferal/go/entry"
	"stferal/go/entry/types/media/video"
	"stferal/go/entry/types/set"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
	p "path/filepath"
)

func ServeFile(s *server.Server, w http.ResponseWriter, r *http.Request, path *paths.Path) {
	err := serveFile(s, w, r, path)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
	}
}

func serveFile(s *server.Server, w http.ResponseWriter, r *http.Request, path *paths.Path) error {
	section := path.Section()
	lang := head.Lang(r.Host)
	tree := s.Trees[section].Local(s.Flags.Local)[lang]
	e, err := tree.LookupEntryHash(path.Hash)
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
	video, ok := e.(*video.Video)
	if ok {
		if p.Ext(path.SubFile.Name) == ".vtt" {
			serveStatic(w, r, video.SubtitleLocation(path.SubFile.Size))
			return nil
		}
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
	set, ok := col.(*set.Set)
	if ok && path.SubFile.Name == "cover.jpg" && set.Cover != nil {
			return serveSingleBlob(w, r, set.Cover, path)
	}
	e, ok := col.(entry.Entry)
	if p.Ext(path.SubFile.Name) == ".vtt" {
		file := p.Join(e.File().Path, path.SubFile.Name)
		serveStatic(w, r, vttPath(file, path.SubFile.Size))
		return nil
	}

	return fmt.Errorf("serveCollectionBlob: File %v not found.", path.SubFile.Name)
}

func vttPath(path, lang string) string {
	l := len(path)
	if l < 5 {
		return path
	}
	return fmt.Sprintf("%v-%v.vtt", path[:l-4], lang)
}

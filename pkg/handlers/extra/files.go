package extra

import (
	"fmt"
	"log"
	"net/http"
	//"path/filepath"
	"stferal/pkg/entry"
	"stferal/pkg/paths"
	"stferal/pkg/server"
)

func Files(s *server.Server, w http.ResponseWriter, r *http.Request, p *paths.Path) {
	// TODO: Panic possible.
	eh, err := s.Trees[p.Page].LookupAcronym(p.Acronym)
	if err != nil {
		s.Debug(err)
		http.NotFound(w, r)
		return
	}

	// What would be an example for this?
	/*
		if p.Type == "files" {
			f, err := entry.ElFileSafe(eh)
			if err != nil {
				s.Debug(err)
				http.NotFound(w, r)
				return
			}
			serveStatic(w, r, filepath.Join(f.Hold.File.Path, p.Descriptor))
			return
		}
	*/

	serveCacheFile(w, r, eh, p.Descriptor)
}

func serveCacheFile(w http.ResponseWriter, r *http.Request, eh interface{}, descriptor string) {
	name, size := paths.SplitDescriptor(descriptor)

	var abs string
	var err error

	switch eh.(type) {
	case *entry.Image:
		abs, err = eh.(*entry.Image).ImageAbs(size), nil
	case *entry.Set:
		abs, err = findSetFile(eh.(*entry.Set), name, size)
	case *entry.Hold:
		abs, err = findHoldFile(eh.(*entry.Hold), name, size)
	default:
		err = fmt.Errorf("Cannot search cache file in #%v#. %v", entry.Type(eh), eh)
	}

	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	serveStatic(w, r, abs)
}

func findHoldFile(h *entry.Hold, name, size string) (string, error) {
	//if name == "cover.jpg" && set.Cover != nil {
	//	return set.Cover.ImageAbs(size), nil
	//}

	for _, e := range h.Els {
		switch e.(type) {
		case *entry.Image:
			if e.(*entry.Image).File.Base() == name {
				return e.(*entry.Image).ImageAbs(size), nil
			}
		}
	}

	return "", fmt.Errorf("Could not find cache file (%v) in Hold (%v)", name, h)
}

func findSetFile(set *entry.Set, name, size string) (string, error) {
	if name == "cover.jpg" && set.Cover != nil {
		return set.Cover.ImageAbs(size), nil
	}

	for _, e := range set.Els {
		switch e.(type) {
		case *entry.Image:
			if e.(*entry.Image).File.Base() == name {
				return e.(*entry.Image).ImageAbs(size), nil
			}
		case *entry.Audio:
			if e.(*entry.Audio).File.Base() == name {
				return e.(*entry.Audio).File.Path, nil
			}
		case *entry.Video:
			if e.(*entry.Video).File.Base() == name {
				return e.(*entry.Video).File.Path, nil
			}
		}
	}

	return "", fmt.Errorf("Could not find cache file (%v) in Set (%v)", name, set)
}

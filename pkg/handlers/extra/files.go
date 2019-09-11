package extra

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"st/pkg/el"
	"st/pkg/paths"
	"st/pkg/server"
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
	if p.Type == "files" {
		f, err := el.ElFileSafe(eh)
		if err != nil {
			s.Debug(err)
			http.NotFound(w, r)
			return
		}
		serveElFile(w, r, f.Path, p.Descriptor)
		return
	}

	serveCacheFile(w, r, eh, p.Descriptor)
}

func serveElFile(w http.ResponseWriter, r *http.Request, path, descriptor string) {
	serveStatic(w, r, filepath.Join(path, descriptor))
}

func serveCacheFile(w http.ResponseWriter, r *http.Request, eh interface{}, descriptor string) {
	_, size := paths.SplitDescriptor(descriptor)

	var abs string
	var err error

	switch eh.(type) {
	case *el.Image:
		abs, err = eh.(*el.Image).ImageAbs(size), nil
		/*
			case *el.Hold:
				abs, err = findCacheFileHold(eh.(*el.Hold), name, size)
			case *el.Set:
				abs, err = findCacheFile(eh.(*el.Set), name, size)
		*/
		/*
			case *el.Hold:
				abs, err = findCacheFileHold(eh.(*el.Hold), name, size)
		*/
	default:
		err = fmt.Errorf("Cannot search cache file in #%v#. %v", el.Type(eh), eh)
	}

	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	serveStatic(w, r, abs)
}

/*
func findCacheFileHold(h *el.Hold, name, size string) (string, error) {
		//if name == "cover.jpg" && set.Cover != nil {
		//	return set.Cover.ImageAbs(size), nil
		//}

	for _, e := range h.Els {
		switch e.(type) {
		case *el.Image:
			if e.(*el.Image).File.Base() == name {
				return e.(*el.Image).ImageAbs(size), nil
			}
		}
	}

	return "", fmt.Errorf("Could not find cache file (%v) in Hold (%v)", name, h)
}

func findCacheFile(set *el.Set, name, size string) (string, error) {
	if name == "cover.jpg" && set.Cover != nil {
		return set.Cover.ImageAbs(size), nil
	}

	for _, e := range set.Els {
		switch e.(type) {
		case *el.Image:
			if e.(*el.Image).File.Base() == name {
				return e.(*el.Image).ImageAbs(size), nil
			}
		}
	}

	return "", fmt.Errorf("Could not find cache file (%v) in Set (%v)", name, set)
}

*/

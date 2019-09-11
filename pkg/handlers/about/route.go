package about

import (
	"net/http"
	"stferal/pkg/el"
	"stferal/pkg/head"
	"stferal/pkg/paths"
	"stferal/pkg/server"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if path == "/about/" || path == "/ueber/" {
		Hold(s, w, r, s.Trees["about"])
		return
	}

	p := paths.Split(path)

	/*
		if p.Type != "" {
			serveFiles(s, w, r, p)
			return
		}
	*/

	hold, err := findHold(s, head.Lang(r.Host), p)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	Hold(s, w, r, hold)
}

func findHold(s *server.Server, lang string, p *paths.Path) (*el.Hold, error) {
	hold, err := s.Trees["about"].Search(p.Name, lang)
	if err == nil {
		return hold, nil
	}

	if p.Acronym != "" {
		hold, err := s.Trees["about"].LookupAcronym(p.Acronym)
		if err == nil {
			return hold.(*el.Hold), nil
		}
		return nil, err
	}

	return nil, err
}

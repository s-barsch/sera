package about

import (
	"net/http"
	//"stferal/go/entry"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	p, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	lang := head.Lang(r.Host)

	rel := p[len("/about"):] // same length as "ueber"

	if rel == "" {
		ServeAbout(s, w, r, s.Trees["about"].Public[lang])
		return
	}
/*

	path := paths.Split(path)

	stru, err := findHold(s, head.Lang(r.Host), p)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ServeAbout(s, w, r, t)
	*/
}

/*
func findHold(s *server.Server, lang string, p *paths.Path) (*entry.Hold, error) {
	hold, err := s.Trees["about"].Search(p.Name, lang)
	if err == nil {
		return hold, nil
	}

	if p.Acronym != "" {
		hold, err := s.Trees["about"].LookupAcronym(p.Acronym)
		if err == nil {
			return hold.(*entry.Hold), nil
		}
		return nil, err
	}

	return nil, err
}
*/
